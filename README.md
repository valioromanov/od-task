# od-task
# Overview
This service provides an API to access detailed information about specific rentals or to retrieve a list of rentals based on their preferences. It exposes two endpoints, allowing clients to retrieve details about individual rentals by ID and retrieve an array of rentals based on specified filters.

# Table of contents

- [Endpoints](#endpoints)
- [Prerequisites](#prerequisites)
- [Running](#running)
- [Unit tests](#tests)

# Endpoints

- GET /rentals/:rentalID
- GET /rentals[?'<filters>']

## Prerequisites
You need to clone <a href="https://github.com/outdoorsy/interview-challenge-backend">interview-challenge-backend</a> first to have a ready-to-use database. After cloning it you can navigate to its folder from the terminal and run the command `docker-compose up`(you need a running
Docker). This will create a container with the database, create the necessary tables, and insert records in them. </br>
</br>
You need to set some env variables:</br>
`- HOST=...` </br>
`- PORT=...` </br>
`- DB_HOST=...` </br>
`- DB_PORT=...` </br>
`- DB_USER=...` </br>
`- DB_PASSWORD=...` </br>
`- DB_NAME=...` </br>

# Running
In this repo, you can find a `.env` file with the necessary values. You can run `source .env` from the parent directory of the project.
Then navigate to `cd cmd` and run `go run main.go` 

# Tests
You need Ginkgo to run unit tests with the make commands. You can get it with `go install github.com/onsi/ginkgo/v2/ginkgo`.If you encounter any issues, ensure that your Go bin directory is in your system's `PATH`( run `export PATH=$PATH:$(go env GOPATH)/bin` if you encounter the problem). To run a make command, you have to be in the parent directory of the project.
You can run the tests with `go test` either.
