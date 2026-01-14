package main

import (
	// "log"
	"gin/gorm/api/config"
	"gin/gorm/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	// router.LoadHTMLGlob("templates/*")
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")

	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatal(err)
	// }

}
