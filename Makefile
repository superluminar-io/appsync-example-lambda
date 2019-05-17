ENV ?= stable
PREFIX ?= $(ENV)
VERSION ?= v0.0.0-dev
PROJECT_NAME := appsync-example-lambda

AWS_PROFILE ?= $(PROJECT_NAME)-$(ENV)
AWS_REGION ?= eu-central-1
AWS_NAME := $(PREFIX)-$(PROJECT_NAME)

PATH_FUNCTIONS := ./src/
LIST_FUNCTIONS := $(subst $(PATH_FUNCTIONS),,$(wildcard $(PATH_FUNCTIONS)*))

clean:
	@ rm -rf ./dist

test:
	@ go test ./...

build-%:
	@ GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build \
		-a -installsuffix cgo -ldflags="-w -s" \
		-o ./dist/$*/handler ./src/$*

build:
	@ $(MAKE) clean
	@ $(MAKE) $(foreach FUNCTION,$(LIST_FUNCTIONS),build-$(FUNCTION))

configure:
	@ aws s3api create-bucket \
		--profile $(AWS_PROFILE) \
		--region $(AWS_REGION) \
		--bucket $(AWS_NAME) \
		--create-bucket-configuration LocationConstraint=$(AWS_REGION)

package:
	@ aws cloudformation package \
		--profile $(AWS_PROFILE) \
		--region $(AWS_REGION) \
		--template-file ./template.yml \
		--s3-bucket $(AWS_NAME) \
		--output-template-file ./dist/stack.yml

deploy:
	@ aws cloudformation deploy \
		--profile $(AWS_PROFILE) \
		--region $(AWS_REGION) \
		--template-file ./dist/stack.yml \
		--capabilities CAPABILITY_IAM \
		--stack-name $(AWS_NAME) \
		--parameter-overrides \
			APIName=$(AWS_NAME) \
			APIKeyExpiration=$(shell echo $$(( $(shell date +%s) + 25920000 )))
