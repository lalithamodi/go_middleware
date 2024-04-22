package main

import (
	"github.com/gin-gonic/gin"
	"main.go/middleware"
)

func main() {
	router := gin.New()

	//router.Use(middleware.Authenticate)  //applying to all

	//admin := router.Group("/admin", middleware.Authenticate)
	//	admin.GET("/getData", getData)
	//admin.GET("/getData1", getData1)
	//}

	router.GET("/getData", middleware.Authenticate, middleware.AddHeader, getData)
	router.GET("/getData1", getData1)

	router.GET("/getData2", getData2)
	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi its get data method",
	})
}

func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi its get data method",
	})
}

func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi its get data method",
	})
}
