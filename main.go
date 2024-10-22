package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
}

var books []Book

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.GET("/books/:id", getBookByID)
    router.POST("/books", createBook)
    router.PUT("/books/:id", updateBook)
    router.DELETE("/books/:id", deleteBook)
    router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, book := range books {
        if book.ID == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func createBook(c *gin.Context) {
    var newBook Book
    if err := c.ShouldBindJSON(&newBook); err == nil {
        books = append(books, newBook)
        c.JSON(http.StatusCreated, newBook)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

func updateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book
    if err := c.ShouldBindJSON(&updatedBook); err == nil {
        for i, book := range books {
            if book.ID == id {
                books[i] = updatedBook
                c.JSON(http.StatusOK, updatedBook)
                return
            }
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
