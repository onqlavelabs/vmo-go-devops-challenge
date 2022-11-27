package main

import "github.com/dinhtp/vmo-go-devops-challenge/application/cmd"

// @title           Programming Challenge Swagger API
// @version         1.0
// @description     Sample Application Swagger API.

// @contact.name   Desmond
// @contact.email  dinhtp@vmodev.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
    cmd.Execute()
}
