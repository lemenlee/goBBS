package main

import (
	"bbs/api"
	"bbs/docs"

	// "bbs/docs"                                   // docs is generated by Swag CLI, you have to import it.
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title BBS Example API
// @version 1.0
// @description This is a sample server BBS server.

// @host lemonlee.net:403
// @BasePath /api/v1

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization

func main() {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	// docs.SwaggerInfo.BasePath = "/v2"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := api.SetupRouter()
	url := ginSwagger.URL("http://lemonlee.net:403/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":403")
}
