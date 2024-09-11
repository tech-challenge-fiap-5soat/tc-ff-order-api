# Fiap Tech Fast Food

## Project Overview

Fiap Tech Fast Food is a system designed to manage a neighborhood fast food restaurant. The system allows for user registration, product management, order creation, and payment processing. It is built to be resilient to failures and scalable.

This project is a specialized management system for the ordering operations of a neighborhood fast food chain. It was developed to solve specific problems related to order processing and management, offering functionalities such as:

- User registration
- Product registration
- Order creation and management
- Payment request processing

The system is built to be fault resilient and scalable, allowing the kitchen to operate efficiently even during periods of high order volume.

## Badges
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-order-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-order-api)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-order-api&metric=bugs)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-order-api)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-order-api&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-order-api)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-order-api&metric=coverage)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-order-api)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-order-api&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-order-api)

To solve a problem of a neighborhood fastfood, a system was created to manage the fastfood, where it is possible to register users, products, create orders and make payments, in addition to the system being resilient to failures and scalable.


### Tech

This api was built using [Golang](https://golang.org/) and some tools:
 * [gin](http://github.com/gin-gonic/gin) - Web framework 
 * [mongo-driver](http://go.mongodb.org/mongo-driver) - driver to deal with MongoDB
 * [viper](https://github.com/spf13/viper) - Config solution tool
 * [mockery](https://github.com/vektra/mockery) - Mock tool to use on unit tests
 * [swag](https://github.com/swaggo/swag) - Tool to generate swagger documentation
 * [docker](https://www.docker.com/) - Containerization tool
 * [docker-compose](https://docs.docker.com/compose/) - Tool to define and run multi-container Docker applications
 * [make](https://www.gnu.org/software/make/) - Tool to define and run tasks
 * [mermaid](https://mermaid-js.github.io/mermaid/#/) - Tool to create diagrams and flowcharts
 * [kubernetes](https://kubernetes.io/pt-br/) - Tool to automate deployment, scaling, and management of containerized applications


## Architecture

For a demonstration of the architecture, visit: [Architecture Video](https://drive.google.com/file/d/1NheE489Ma2W28Jvz3ZzRNAWCeHTrwVbm/view?usp=sharing)

## Running the Application

### Using Docker

The app can be started using docker and you can use the actions pre-defineds on Makefile

* ***Build image***

To build an image from project to push to a registry you can use the command below:

```sh
make build-image
```
this command will generate an image with this tag: `tc-ff-order-api`

#### Generate Documentation

To generate the documentation to publish on the project like an OpenAPI, use the command below:

```sh
make serve-swagger
```
this command will generate a directory called `docs` 

### Development

Before run the application, you need to export the variables below:

```sh
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export MONGODB_HOST=localhost
export MONGODB_PORT=27017
export MONGODB_DATABASE=db
export MONGODB_USER=root
export MONGODB_PASS=root
```

To run in development for debug or improvement you can use another command:

```sh
make start-local-development &
```

And run that command in another terminal:

```sh
make run
``` 

This command will start a container with hot-reload to any modification on the code. Including a container with an instance of MongoDB.

To stop the container execute:

```sh
make stop-local-development
```

### Test

Locally you can use the command below:

```sh
go test ./...  -v
```

or use a make action: 

```sh
make test   
```

## Configuration

Configuration settings are managed using environment variables and a configuration file. Refer to the `configs.yaml.sample` file for the required settings.

```yaml:src/external/api/infra/config/configs.yaml.sample
startLine: 1
endLine: 20
```
