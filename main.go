package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title" binding:"required"`
	Author string  `json:"author" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type BookHandler struct {
	DB *sql.DB
}

func main() {

	connStr := "user=postgres password=example dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()


	err = db.Ping()
	if err != nil {
		fmt.Println("Erro: não consegui me conectar ao db!")
		panic(err)
	}
	fmt.Println("Sucesso, conectado ao db!")


	handler := &BookHandler{DB: db}


	router := gin.Default()
	

	router.POST("/books", handler.postBooks)


	fmt.Println("Servidor rodando em http://localhost:8080")
	router.Run("localhost:8080")
}

func (h *BookHandler) postBooks(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido: " + err.Error()})
		return
	}


	query := `INSERT INTO books (title, author, price) VALUES ($1, $2, $3) RETURNING id`

	err := h.DB.QueryRow(query, newBook.Title, newBook.Author, newBook.Price).Scan(&newBook.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível salvar no banco: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}
