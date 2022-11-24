## About
Go programming challenge for `onqlavelabs`. 
This is an HTTP gateway for basic management of `application` entity with a simple JWT authentication. 


## Prerequisite:
- Linux based OS is preferred
- Docker (version 20.10.21 is preferred) with compose plugin or *docker-compose*


## Spin-up Project
- Serve HTTP gateway: `docker compose up --build`
- API Docs: `http://localhost:8080/swagger/index.html`
- Default username and password: `root/root`


## Technical choice:
- HTTP gateway: [Golang Echo](https://echo.labstack.com/)
- MongoDB: [Go Driver](https://www.mongodb.com/docs/drivers/go/current/#mongodb-go-driver)
- API Doc Swagger: [swaggo](https://github.com/swaggo/swag), [echo-swagger](https://github.com/swaggo/echo-swagger)
- API Authentication: [Golang JWT](https://github.com/golang-jwt/jwt)
- Customer logger: [Logrus](github.com/sirupsen/logrus)
- Project command and config: [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper)


## Architecture
### Project Architecture
- Domain based approach with a loosely applied hexagonal architecture is utilized for this project.
- `cmd` directory is the default directory created by `Cobra` for a command based project. This also includes the migration, serve and seed command.
- The core or `domain` logic is located within the `domain` directory.
- Within the `domain` directory, each entity has its own directory containing its specific logic and handler.
- `controller`, `database` and`server` directories are external layers wrapping around the `domain` layer.
- `model` and `message` directories define the http message payload and MongoDB model representation.
- utility directory is listed as: `logger`, `docs` (swagger API docs), `util`

### Project Structure
```
├── cmd
│ ├── migrate.go
│ ├── migration.go
│ ├── root.go
│ ├── seed.go
│ └── serve.go
├── controller
│ ├── application.go
│ ├── auth.go
│ └── swagger.go
├── database
│ ├── constant.go
│ ├── json_schema.go
│ ├── migration.go
│ └── mongodb.go
├── db.json
├── docker-compose.yaml
├── Dockerfile
├── docs
│ ├── docs.go
│ ├── swagger.json
│ └── swagger.yaml##
├── domain
│ ├── application
│ │ ├── data.go
│ │ ├── repository.go
│ │ ├── service.go
│ │ └── service_test.go
│ └── auth
│     ├── constant.go
│     ├── service.go
│     └── service_test.go
├── go.mod
├── go.sum
├── LICENSE
├── logger
│ └── log.go
├── main.go
├── Makefile
├── message
│ ├── application.go
│ ├── auth.go
│ └── constant.go
├── model
│ ├── application.go
│ └── model.go
├── Results.md
├── server
│ └── echo.go
└── util
    ├── converter.go
    └── file.go
```

### Assessments
* **Architecture**: the logic is divided into 3 separate layer: gateway, service, and model layer. Each layer has one single purpose (gateway to handle API call, service for data handling and conversion, model for database transaction) and independent of one and another.
* **Clarity**: The project purpose, architecture, structure and assessment criteria are explained clearly in the `Results.md` file.
* **Correctness**: basic management features (CRUD) are met with additional filter options in listing API.
* **Code quality**: The source code is well organized with few nested functions and directory along with clear structure, logic block and easy to navigate structure.
* **Usability**: the project is suitable for multiple clients, even with physical constraint as the source code is relatively lightweight and required no extensive CPU or memory to execute.
* **Security**: No obvious security flaw is witness during the development and testing process.
* **Testing**: basic unit test for some important functions is included to demonstrate testing coverage of the project.
* **Technical choices**: the decision of using Echo for HTTP gateway and Golang Mongo Driver are both lightweight and suitable for this kind of project given the time-constraint and project complexity.
* **Documentation**: All API endpoints are documented and presented in a swagger documentation.
* **Production-readiness**: error handling and logging are included. Monitoring and tracing are not implemented due to the time-constraint. However, this aspect can be approached with 
* **Scalability**: The scalability can be achieved relative easily thanks to the separation of 3 important layers.
* **Authentication**: simple authentication with JWT bearer token is implemented for authentication. The process is not directly point into a database for user validation since this aspect mostly focus on the authentication part
