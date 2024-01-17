# Running tests against local PostgreSQL database

To run the tests against the local database, make the following environment setup and configuration changes before running the tests:

1. Set up the local postgres database using postgres docker image:
```bash
docker run --name postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 postgres
```
2. Set environment variables to let the test connect to the database:
```bash
export PGUSER="postgres"
export PGPASSWORD="your_password"
export PGDATABASE="postgres"
export PGSSLMODE="disable"
```
3. Run the tests:
```bash
go test -v
```

To run the tests locally, you have to spin up a local postgres database, and set environment
variables to let the test connect to the database.

## Get postgres docker image
cd into your project root directory and run the following command to install the required dependencies:
```
 go get github.com/DATA-DOG/go-sqlmock
```
`docker pull postgres
```

## Run postgres image and forward port to localhost
`docker run --name postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 postgres`

## Run tests
`env PGUSER="postgres" PGPASSWORD="your_password" PGSSLMODE="disable" go test -v`
