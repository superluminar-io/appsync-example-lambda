# AppSync Example w/ Lambda

[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](http://github.com/superluminar-io/appsync-example-lambda/blob/master/LICENSE.md)

> Deploy a GraphQL API using AWS AppSync, Serverless Application Model, and AWS Lambda functions using Go.

## Usage

```bash
# Create S3 Bucket for CloudFormation Artifacts
$ > AWS_PROFILE=your-profile-name make configure

# Build, Package, and Deploy the CloudFormation Stack
$ > AWS_PROFILE=your-profile-name make build package deploy
```

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
