package main

import (
	"munayfund-api2/docs"
	"munayfund-api2/internal/infra/restapi"
)

// @title           Swagger MunayFund API WEB 2
// @version         1.0
// @description     This API is in charge of leveraging some functionalities of munayFund a next gen crowdfunding solution.
// @termsOfService  https://swagger.io/terms/

// @contact.name   KaniKrafters
// @contact.url    https://todefine.github.com
// @contact.email  todefine@github.com

// @license.name  Apache 2.0
// @license.url   https://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	docs.SwaggerInfo.BasePath = "/api/munayfund/v1"
	restapi.NewAPI().StartAPI()
}
