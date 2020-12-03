package service

import (
	"time"

	"Walter0697/GinBackend/base"
	"Walter0697/GinBackend/apibody"
	"Walter0697/GinBackend/models"
)

func GetAllBooks() []models.Book {
	var books []models.Book
	base.DB.Where("is_deleted = false").Find(&books)
	return books
}

func CreateBook(input apibody.CreateBookInput, userId uint) models.Book {
	book := models.Book{Title: input.Title, Author: input.Author, LastModifiedTime: time.Now(), LastModifiedUID: userId, CreateTime: time.Now(), CreateUID: userId, IsDeleted: false }
	base.DB.Create(&book)
	return book
}

func GetBookById(bookId uint) (*models.Book, error) {
	var book models.Book
	if err := base.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func UpdateBook(bookId uint, input apibody.UpdateBookInput, userId uint) (*models.Book, error) {
	var book models.Book

	if err := base.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		return nil, err
	}

	base.DB.Model(&book).Updates(map[string]interface{}{"title": input.Title, "author": input.Author, "last_modified_time": time.Now(), "last_modified_uid": userId})
	return &book, nil
}

func DeleteBook(bookId uint, userId uint) (bool, error) {
	var book models.Book
	if err := base.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		return false, err
	}

	base.DB.Model(&book).Updates(map[string]interface{}{"is_deleted": true});
	
	//instead of deleting the book, we set isDeleted to true so that we can still see the deleted records
	//base.DB.Delete(&book)

	return true, nil
}