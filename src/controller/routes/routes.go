package routes

import (
	"github.com/aristidesneto/crud-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	router.GET("/user/:id", controller.FindUserByID)
	router.GET("/user/email/:email", controller.FindUserByEmail)
	router.POST("/user", controller.StoreUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
}
