
# Inventapp_v3

This project is an API for an Inventory Management System (IMS) application. It is implemented in Go and provides RESTful APIs and gRPC APIs to interact with the system.
 

## Installation

To run this project, you need to have Go installed on your system. You can download and install it from the [official Go website](https://golang.org/).

  

Once you have Go installed, you can clone this repository:

  

```bash
git  clone  https://github.com/nicodanke/inventapp_v3.git
```

  
Then you have to install:
- GNU make: to run the scripts necessary to run the project. You can get more information in the [official web site](https://www.gnu.org/software/make/).
- migrateCLI  to run the DB migrations. You can download it from the [official repository](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). In Mac use the following command to install:
```bash
brew  install  golang-migrate
```
- SQLC to generate SQL queries and models. You can get more information about SQLC in its [documentation](https://docs.sqlc.dev/en/latest/index.html). To install SQLC in Mac just run:
```bash
brew install sqlc
```
- DBDocs to automatically generate DB documentation. To install it and lear how to use it you can go to the [official documentation](https://dbdocs.io/docs). To visualize DBML files properly in VSCode, it is recommended to install [DBML Language extension](https://marketplace.visualstudio.com/items?itemName=duynvu.dbml-language). In order to install DBDocs just run:

```bash
npm install -g dbdocs
```
- DBML CLI is used to generate SQL Schema using the db.dbml file. To install it, just run:

```bash
npm install -g @dbml/cli
```

## Configuration

For the project to run, you have to set the right configuration in your project, for this reason, the first step is to check the `.env` file in the root project, where the configuration of the DB is placed. The .env file is only used in the Makefile to run the DB container and the migrations.

To check the configuration that the app is going to use, you should check the `.env` file. Here you can configure among other things the IP and port where the RESTful and gRPC APIs will be running.

The first time you run the project you should create a container for the database with the following command:
```bash
make create_container
```

Then you should create the database in the container with the following command:
```bash
make create_db
```

If it is not the first time you run the project but the container of the DB is not up, then you should only run this command:
```bash
make start_container
```

To stop all container running, you can run:
```bash
make stop_containers
```

To configure DBDocs, in order to generate new documentatio, you just need to run:
```bash
dbdocs login
```

This will ask you how to authenticate, and you have to choose email option. Then an email will be sent to the address of the owner of the project with an OTP code that you must enter in the terminal to complete the authentication process of DBDocs.

## Usage

  

### Running the Project

To run the project, navigate to the project directory and use the following command:
```bash
make server
```

This will start the server, and it will be accessible via RESTful APIs and gRPC APIs. By default, RestAPI will run on `localhost:8080` and gRPC API will run on `localhost:9090`.  This configuration can be changed in `app.env` file.

  
### Running Tests

We have included unit tests to ensure the reliability of the codebase. You can run the tests using the following command:
```bash
make test
```

## How to

### Create New Migration
In order to create a new migration, you just have to run:
```bash
migrate create -ext sql -dir db/migrations -seq migration_name
```
Replacing `migration_name` with the proper name of the migration. This will create two files in the `db/migrations` folder that correspond to the migration up file and the migration down file. Add the migration up and down code to the migration files created and save both files. To write the migration code you can add the table structure in the `doc/db.dbml` file, then run:
```bash
make db_schema
```
to create the doc/scheme.sql file, where all the DDL of the database will be placed. You can extract the piece of code that adds the new schemas and add them inside the migration files.

The migration will be applied automatically when you run the project, but if you want to apply all the migrations manually you can run:
```bash
make migrate_up
```

To remove all the migrations applied manually, you can run:
```bash
make migrate_down
```

To only apply the last migration manually you should run:
```bash
make migrate_up1
```

To remove the last migration applied manually, you can run:
```bash
make migrate_down1
```

### Add new queries

In order to add a new query you just need to add the query inside `db/query` in a `.sql` file. To add the query you just have to follow the [SQLC documentation](https://docs.sqlc.dev/en/latest/howto/select.html#). After adding the necessary queries, you just have to run
```bash
make sqlc_generate
```

to generate the SQL queries and models. The models and queries are going to be generated inside `db/sqlc` folder.

### Generate DB Documentation
To generate DB documentation you just have to add the DB structure inside `doc/db.dbml` file. And then run:
```bash
make db_docs
```
The documentation will be available in [this page](https://dbdocs.io/nicodanki/InventApp). The password to access is: *inventappDoc2024*


## Contributing

If you'd like to contribute to this project, please follow these guidelines:

1. Create a new branch for your feature or bug fix.

2. Make your changes and ensure that the code passes all tests.

3. Submit a pull request detailing your changes.

  

## Best Practices

To maintain consistency and readability in the codebase, please adhere to the following best practices:

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).

- Write clear and concise commit messages.

- Keep the codebase clean and well-documented.

- Follow the project's coding style and conventions.

  

## License

This project is licensed under the [MIT License](LICENSE).