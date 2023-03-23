package routes

// API route handler folder to manage all the user-related routes in our application

import (
	"gin-mongo-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())

	router.GET("/user/:userId", controllers.GetAUser())

	router.PUT("/user/:userId", controllers.EditAUser()) //add this

	router.DELETE("/user/:userId", controllers.DeleteAUser()) //add this

	router.GET("/users", controllers.GetAllUsers()) //add this

}
