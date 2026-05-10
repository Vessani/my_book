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
	// 1. Configuração da Conexão
	connStr := "user=postgres password=example dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// O defer garante que o banco só feche quando o SERVIDOR parar
	defer db.Close()

	// 2. Verificação do Banco
	err = db.Ping()
	if err != nil {
		fmt.Println("Erro: não consegui me conectar ao db!")
		panic(err)
	}
	fmt.Println("Sucesso, conectado ao db!")

	// 3. Inicialização do Handler com a Injeção de Dependência
	handler := &BookHandler{DB: db}

	// 4. Configuração das Rotas
	router := gin.Default()
	
	// Registro da rota POST
	router.POST("/books", handler.postBooks)

	// 5. Execução do Servidor (O programa "trava" aqui e fica esperando o curl)
	fmt.Println("Servidor rodando em http://localhost:8080")
	router.Run("localhost:8080")
}

func (h *BookHandler) postBooks(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido: " + err.Error()})
		return
	}

	// Dica: Certifique-se que o nome da tabela no DB é "books" (minúsculo é o padrão SQL)
	query := `INSERT INTO books (title, author, price) VALUES ($1, $2, $3) RETURNING id`

	err := h.DB.QueryRow(query, newBook.Title, newBook.Author, newBook.Price).Scan(&newBook.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível salvar no banco: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}
