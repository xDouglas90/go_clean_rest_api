## Golang REST API - Clean Architecture Principles

## Install Mux library

```bash
go get -u github.com/gorilla/mux
```

## Install Chi library

```bash
go get -u github.com/go-chi/chi
```

## Install Firestore library

```bash
go get -u cloud.google.com/go/firestore
```

## Install MySQL library

```bash
go get -u github.com/go-sql-driver/mysql
```

## Install MongoDB library

```bash
go get -u go.mongodb.org/mongo-driver
```

## Install Testify library

```bash
go get github.com/stretchr/testify
```

## Export Environment variable 

```bash
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private-key.json'
```

## How to get the private key JSON file:
## From the Firebase Console: Project Overview -> Project Settings -> Service Accounts -> Generate new private key

## Build

```bash
go build
```
## Test (specific test)

```bash
go test -run NameOfTest
```

## Test (all the tests within the service folder)

```bash
go test service/*.go
```

## Run

```bash
go run .
```

```bash
go run *.go
```
