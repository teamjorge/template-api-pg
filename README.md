# template-api-pg

    A quickstart template for building a golang API with Postgres integration.

## Getting Started

* Clone this repo or use the Github `Use This Template` functionality.
* Delete the `.git` directory and run `git init`.
* Rename the parent directory and `go.mod` module to your project's name

*If using the VSCode DevContainer:*

* Update the `.devcontainer/.env` to match your needs
* Open as a VSCode DevContainer

*If you plan to run it without a .devcontainer:*

* Update the `config.yaml` to match your needs. *This is mostly done because the DevContainer easily manages environment variables for you, but you prefer to use environment variables you can use the .devcontainer/.env as a base*
* Ensure that you have installed both sqlboiler and sql-migrate:
  * `go install github.com/volatiletech/sqlboiler/v4@latest`
  * `go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest`
  * `go install github.com/rubenv/sql-migrate/...@latest`

You can now apply migrations:

```shell
cd sql
sql-migrate up
```

Build and run the server
```shell
go build -o . ./...
./server
```

## Features

### Niceties

* DevContainer with everything pre-installed
* Server and Database connections pre-configured
* Out-of-the-box middleware and helper endpoints
* Config already plugged in and easily configured

### Third-party Libraries

* sqlboiler - "Database-first" ORM
* sqlmigrate - SQL Schema migration tool
* gorilla/mux - Routing and API enhancement
* negroni - Middleware
* spf/viper - Config management

## Specifics

### Config

Configuration is handled by [`viper`](https://github.com/spf13/viper) and is recommended to be driven by environment variables. A config file has been supplied for cases where using a DevContainer or environment variables will be troublesome/impossible.

The `internal/config/init.go` contains the initialisation code for the config properties. This most likely is not the most effective of doing it, so feel free to change it.

The config package is by default initialised in the `cmd/server/main.go` with the `_ "template-api-pg/internal/config"` import. This ensures that the `init` function is called and viper has pre-parsed any variables specified. Should you wish to create an additional executable that will utilise the config files, please include the same import in that `main.go` file.

### Database

Migrations can be configured via the config found in `sql/dbconfig.yml` and added to the `sql/migrations` directory.

The following commands need to be run from the `sql` directory.

**Applying all migrations**: `sql-migrate up`

**Rolling back all migrations**: `sql-migrate down`

Please see the [sql-migrate repo](https://github.com/rubenv/sql-migrate) for additional information.


### Models

Generating the models happens via the `sqlboiler` library. The configuration file can be found in `sql/sqlboiler.toml`.

For additional configuration options, please see the [sqlboiler repo](https://github.com/volatiletech/sqlboiler).

To generate the models, please run `sqlboiler psql` from the `sql` directory. Model files will be wiped and re-generated to the `models` directory.

### API

Having a look at the Example controller is the best way to delve into the API. The Example model and REST endpoints have been created to show the general flow of the API. This is a very simple CRUD style endpoint, but the general flow will remain the same for more complex endpoints.

You are able to test the various endpoints by starting the server and utilising the following requests:

* `OPTIONS /example` - To get the schema
* `POST /example` - Fill in the schema from the OPTIONS request and add as a `json` body
* `GET /example` - Lists all examples
* `GET /example/1` - Retrieves the example with ID `1`
* `PUT /example/1` - Updates the example with ID `1` with the given `json` body. (Only the fields you want to update)
* `DELETE /example/1` - Deletes the example with ID `1`

To remove the Example code:

* `sql-migrate down` (If you have already run the migrations)
* Remove the references the the `example` table in `sql/migrations/0000_initial.sql`
* `sql-migrate up`
* Regenerate the models: `sqlboiler psql`
* Remove the `internal/api/example.go` file
* Remove the line adding the controller in the `addRouting` found in the `internal/api/server.go` file

## What's Next

* Convenience script to handle clones, renaming, and any other additional setup tasks
* Unit tests for the endpoints (using SQL Mocking)
