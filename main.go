package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/lib/pq"
)


type book struct {
	ID int `json:"id"`
	Title string `json:"title" binding:"required"`
	Author string `json:"author" bindgin:"required"`
	Price float64 `json:"price" binding:"required"`
}


/* func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}
*/

func main() {
	connStr := "user=postgres password=example dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	err = db.Ping()

	if err != nil {
		fmt.Println("Erro não consegui me conectar ao db!")
		panic(err)
	}

	fmt.Println("Sucesso, conectado ao db!")
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
