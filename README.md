# Online Store Service

A simple Go service using Gorilla Mux.

## Table of Contents
- [Introduction](#introduccion)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the project](#running-the-project)
- [Documentation](#documentation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction
This project is a simple online store service implemented in Go, utilizing Gorilla Mux for routing.

## Prerequisites
- Go 1.16 or higher

## Installation
1. Clone de repository
2. Navigate to the project directory
3. Initialize the Go module:
`go mod init github.com/aliasthewho/online-store`

## Running the project
1. Run the service
`go run main.go`

2. Run the tests:
`go test`

## Documentation
1. Install swag for generating documentation:
`go install github.com/swaggo/swag/cmd/swag@latest`

2. Get Swagger dependencies:
`go get -u github.com/swaggo/http-swagger`
`go get -u github.com/swaggo/files`

3. Generate the documentation:
`swag init -g cmd/online-store/main.go`

## Usage
Follow instructions

## 1. Creating a project
a. `go mod init github.com/user/project` <br>
b. Define the directory structure as suggested in this example

## 2. Runnig the project
a. Run the service: `go run main.go` <br>
b. Run the test: `go test`

## 3. Docummenting configuration
a. `go install github.com/swaggo/swag/cmd/swag@latest` <br>
b. `go get -u github.com/swaggo/http-swagger` <br>
c. `go get -u github.com/swaggo/files`

## 4. Docummenting execution
a. `swag init -g cmd/online-store/main.go` 
