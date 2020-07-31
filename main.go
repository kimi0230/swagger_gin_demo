package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"swagger_gin_demo/routes"

	"github.com/gin-gonic/gin"

	_ "swagger_gin_demo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/joho/godotenv"
)

/*
@title Swagger Example API With Gin
@version 1.0
@description This is a sample server celler server.
@termsOfService http://swagger.io/terms/

@contact.name API Support
@contact.url http://www.swagger.io/support
@contact.email support@swagger.io

@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html

@host localhost:5566
@BasePath /api/v1
@query.collection.format multi
@x-extension-openapi {"example": "value on a json format"}
*/
func main() {

	if len(os.Args) > 1 {
		env := os.Args[1]
		switch env {
		case "app":
			gin.SetMode(gin.ReleaseMode)
		case "dev":
			gin.SetMode(gin.DebugMode)
		case "qa":
			gin.SetMode(gin.TestMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Listen and Server
	url := "localhost" // os.Getenv("APP_URL")
	port := "5566"     // os.Getenv("APP_PORT")
	serverURL := url + ":" + port

	// import router path
	r := routes.SetupRouter()

	// enable swagger API doc
	if mode := gin.Mode(); mode == gin.DebugMode {
		swagURL := ginSwagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", serverURL))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagURL))
	}
	log.Println("----- start server ----", serverURL)
	log.Fatal(http.ListenAndServe(serverURL, r))
}
