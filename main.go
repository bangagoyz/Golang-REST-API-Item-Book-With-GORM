package main

import (
	"chapter2_4/controller"
	"chapter2_4/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	r := gin.Default()

	r.POST("/books", controller.CreateBook)
	r.GET("/books", controller.GetBooks)
	r.GET("/books/:id", controller.GetBook)
	r.PUT("/books/:id", controller.UpdateBooks)
	r.DELETE("/books/:id", controller.DeleteBook)
	log.Fatal(r.Run(":8080"))

}
