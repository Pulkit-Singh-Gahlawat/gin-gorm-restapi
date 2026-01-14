package routes

import (
	"gin/gorm/api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", nil)
	// })
	
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}

