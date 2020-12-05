package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "Walter0697/GinBackend/service"
  "Walter0697/GinBackend/apibody"
  "Walter0697/GinBackend/helper"
  "Walter0697/GinBackend/utility"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
  books := service.GetAllBooks()

  c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input apibody.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	// Get user Information
	user, _ := helper.GetUserByContext(c)
  
	// Create book
	book := service.CreateBook(input, user.ID)
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) { 
	// Validate input type
	id, err := utility.ConvertStringToUint(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be number"})
	}

	// Find the book from database
	book, err := service.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Validate input
	var input apibody.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	// Validate input type
	id, err := utility.ConvertStringToUint(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be number"})
	}

	// Get user Information
	user, _ := helper.GetUserByContext(c)

	// Update the book
	book, err := service.UpdateBook(id, input, user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
  }

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Validate input type
	id, err := utility.ConvertStringToUint(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be number"})
	}

	// Get user Information
	user, _ := helper.GetUserByContext(c)

	// Delete the book
	isDeleted, err := service.DeleteBook(id, user.ID)
	if isDeleted == false && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }


func Testing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "CI/CD works!"})
}