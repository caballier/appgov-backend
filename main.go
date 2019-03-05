package main

import (
	"net/http"

	controller "github.com/caballier/appgov-backend/controllers"
	"github.com/caballier/appgov-backend/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @contact.name Caio Medeiros Pinto
// @contact.url
// @contact.email caio.medeiros.pinto@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func setupRouter() *gin.Engine {
	route := gin.Default()

	docs.SwaggerInfo.Title = "Applications API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "API to manage applications."

	ctrl := controller.NewController()

	route.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	v1 := route.Group("/api/v1")
	{
		application := v1.Group("/application")
		{
			application.GET("", ctrl.ListApplications)
			application.GET(":id", ctrl.ShowApplication)
			application.POST("", ctrl.AddApplication)
			application.PUT(":id", ctrl.UpdateApplication)
			application.DELETE("", ctrl.DeleteApplication)
		}
	}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
