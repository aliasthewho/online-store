# A simple service with Go

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
