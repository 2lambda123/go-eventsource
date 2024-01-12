# Running tests towards local PostgreSQL database

# Running tests towards local database

To run the tests locally, you have to spin up a local PostgreSQL database and set the environment variables to allow the tests to connect to the database.

## Get PostgreSQL docker image
`docker pull postgres`

## Run postgres image and forward port to localhost
`docker run --name postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 postgres`

## Run tests
`env PGUSER="your_username" PGPASSWORD="your_password" PGSSLMODE="disable" go test -v`
