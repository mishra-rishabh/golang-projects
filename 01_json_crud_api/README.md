# Golang CRUD API

This repository contains a simple Golang implementation of a CRUD (Create, Read, Update, Delete) API. The API is designed to handle basic operations on a collection of resources, demonstrating the fundamentals of building RESTful APIs in Golang.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Important Commmands](#important-commands)

## Features

- Create new records
- Read existing records
- Update records
- Delete records


## Prerequisites

Before you begin, ensure you have the following packages and tools installed:

- [Gin Framework](https://golang.org/dl/)
- [GORM](https://gorm.io/docs/index.html)
- [Postgres](https://gorm.io/docs/index.html)
- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
- [GoDotenv](https://github.com/joho/godotenv)
- Tools
    - **[Elephant SQL](https://www.elephantsql.com/):** To manage database. <br/>
        Signup here and get a connection URL by creating a new instance.

    - **[Table Plus](https://tableplus.com/):** GUI tool to create, query, access database. <br/>
        Connect using the above created URL by importing it in table plus.
        

## Usage

Once the server is running, you can interact with the API using your preferred HTTP client or tools like [Postman](https://www.postman.com/).


## Endpoints

- **GET /items**: Retrieve all items
- **GET /items/{id}**: Retrieve a specific item by ID
- **POST /items**: Create a new item
- **PUT /items/{id}**: Update an existing item by ID
- **DELETE /items/{id}**: Delete an item by ID


## Important Commands

* Use this commands once CompileDaemon is installed.
    ```
    go install github.com/githubnemo/CompileDaemon
    ```

* Run CompileDaemon with your module name (Look for it in `go.mod` file).
    ```
    CompileDaemon -command="./json_crud_api"
    ```

* Run this command once `migrate.go` file is completed.
    ```
    go run migrate/migrate.go
    ```