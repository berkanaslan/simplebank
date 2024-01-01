# Simple Bank

This is a simple bank project that I've created to learn Go.

## Lecture 1: Database

First, I created postgres database named "simple_bank".
Then I've installed a package named migrate.

```shell
brew install golang-migrate
```

To initialize migration folder, run the following command:

```shell
migrate create -ext sql -dir db/migration -seq init_schema
```

Second I've installed the package named sqlc.
It generates Go code from SQL. It helps us to write type-safe SQL queries without having to worry about mundane tasks
like creating and maintaining DTOs or parsing query results.

```shell
brew install sqlc
sqlc init # For the configuration, please refer to the sqlc.yaml file.
```

All tool commands are in the Makefile. Please refer to the Makefile for more details.

## Lecture 2: Database Testing

I've installed the package named testify. It provides a set of packages that help us write code that is easy to test and
CRUD operations of Account, Entry, and Transfer models are using a package named "require" thanks to the Testify.

```shell
go get github.com/stretchr/testify
```

## Lecture 3: Database Transaction

SQLC does not provide us with transaction support with code generation. (Actually it generates WithTx function. I'll
check it later.) So I've implemented the transaction support manually.

### ACID

As a back-end developer, we should know the ACID properties of a transaction.

- Atomicity: All or nothing (either all of the operations are executed or none of them is).
- Consistency: The database should be in a consistent state before and after the transaction.
- Isolation: The intermediate state of the transaction is invisible to other transactions.
- Durability: The changes of a transaction must be permanent.

