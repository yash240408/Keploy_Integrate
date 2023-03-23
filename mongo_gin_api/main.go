package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"github.com/gin-gonic/gin"

	"github.com/keploy/go-sdk/integrations/kgin/v1"
	"github.com/keploy/go-sdk/keploy"
	// "github.com/keploy/go-sdk/integrations/kmongo"   used in configs/setup.go where we connect to the database
)

func main() {
	r := gin.New() // use gin.New() to create a new Gin router

	//initialize Keploy
	port := "6000"
	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "my_app",
			Port: port,
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:6789/",
		},
	})
	kgin.GinV1(k, r) // use Keploy's GinV1 function to initialize the router

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(r) // pass the router to UserRoute

	r.Run(":" + port) // use r instead of router
}

func main() {
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "data": "Hello from Gin-gonic & mongoDB",
    })
  })

	router.Run("localhost:6000")
}
