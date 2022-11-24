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