package main

import (
	"github.com/EikoNakashima/test-go4/controller"
	"github.com/EikoNakashima/test-go4/middleware"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}
	engine.Run(":3000")
}
