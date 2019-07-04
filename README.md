# AppSync Example w/ Lambda

[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](http://github.com/superluminar-io/appsync-example-lambda/blob/master/LICENSE.md)
[![Read Tutorial](https://badgen.now.sh/badge/Read/Tutorial/orange)](https://superluminar.io/2019/07/02/graphql-mit-appsync-und-lambda/)

> Deploy a GraphQL API using AWS AppSync, Serverless Application Model, and AWS Lambda functions using Go. Read more about the implementation details in our [blog post about AppSync and GraphQL](https://superluminar.io/2019/07/02/graphql-mit-appsync-und-lambda/).




## Schema

```graphql
type Person {
	id: Int!
	name: String!
	age: Int!
	birthday: String!

	friends: [Person!]!
}

type Query {
	people: [Person!]!
	person(id: Int): Person!
}

schema {
	query: Query
}
```

## Usage

### Deployment

```bash
# Create S3 Bucket for CloudFormation Artifacts
$ > AWS_PROFILE=your-profile-name \
    make configure

# Build, Package, and Deploy the CloudFormation Stack
$ > AWS_PROFILE=your-profile-name \
    make build package deploy
```

### API Access

```bash
# Print GraphQL API Endoint
$ > AWS_PROFILE=your-profile-name \
    make outputs-GraphQL

https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql

# Print AppSync API Key
$ > AWS_PROFILE=your-profile-name \
    make outputs-APIKey

da2-1jdf4nmbwpsdr4vfxcxfza
```

### Example

```bash
$ > curl \
    -XPOST https://tdk6mhrty7ii.appsync-api.eu-central-1.amazonaws.com/graphql \
    -H "Content-Type:application/graphql" \
    -H "x-api-key:da2-1jdf4nmbwpsdr4vfxcxfza" \
    -d '{ "query": "query { people { name } }" }' | jq
```

## Resolvers

* `Query.people`
* `Query.person`
* `Field.person.age`
* `Field.person.friends`

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.
