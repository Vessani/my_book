package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

)


type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price float64 `json:"price"`
}


func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}

var books = []book{
	{ID: "1", Title: "The Fellowship of the Ring", Author: "J.J Tolkien", Price: 25.99},
	{ID:"2", Title: "The Name of the Wind", Author: "Patrick Rothfuss", Price: 15.50},
	{ID: "3", Title: "Animal Farm", Author: "George Orwell", Price: 13.00},
}

func getBooks (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks (c *gin.Context) {
	var newBook book 

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
