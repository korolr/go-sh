## Stacks

- Go
- GraphQL : [graphql-go](https://github.com/graph-gophers/graphql-go)
- ORM : [gorm](https://github.com/jinzhu/gorm)

## Features

- User Sign Up & Sign In
- Change a Password, Profile

## How to Run

### Initialize DB

1. Create a database

```shell
postgres=# CREATE DATABASE go;
```

2. Create a user as owner of database

```shell
postgres=# CREATE USER go WITH ENCRYPTED PASSWORD 'go';

postgres=# ALTER DATABASE go OWNER TO go;
```

3. Grant all privileges to user for the database

```shell
postgres=# GRANT ALL PRIVILEGES ON DATABASE go TO go;
```

4. Configure the db in `db.go`

```go
// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=go dbname=go password=go sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
```

or with Docker

> host address should be edited to `host.docker.internal` to connect a host interface.

```go
// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {
	db, err := gorm.Open("postgres", "host=host.docker.internal port=5432 user=go dbname=go password=go sslmode=disable")

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}
```

### Initial Migration

```shell
$ go run ./migrations/init.go
```

or with Docker

```
$ docker build -t go-graphql-api .
$ docker run --rm go-graphql-api migrate
```

This will generate the `users` table in the database as per the User Model declared in `./model/user.go`

### Run the server

```shell
$ go run server.go
```

or with Docker

```
$ docker run --rm -d -p 8080:8080 go-graphql-api
```

### GraphQL Playground

Connect to http://localhost:8080

### Authentication : JWT

You need to set the Http request headers `Authorization`: `{JWT_token}`
