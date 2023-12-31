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

TODO