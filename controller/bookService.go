package controller

import (
	"chapter2_4/database"
	"chapter2_4/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookUpdate struct {
	BookName string `json:"name_book"`
	Author   string `json:"author"`
}

func CreateBook(c *gin.Context) {
	db := database.GetDB()

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	db, err := database.StartDB()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, books)

}
func GetBook(c *gin.Context) {
	var book models.Book

	db, err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id= ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBooks(c *gin.Context) {
	var book models.Book

	db, err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	var updateBooks BookUpdate
	if err := c.ShouldBindJSON(&updateBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Model(&book).Updates(models.Book{BookName: updateBooks.BookName, Author: updateBooks.Author}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	db, err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted succesfully",
	})
}
