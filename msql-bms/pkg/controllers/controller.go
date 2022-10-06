package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"1.19/pkg/models"
	"1.19/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
	}
	book := models.GetBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	err := utils.ParseBody(r, CreateBook)
	if err != nil {
		fmt.Println(err)
	}
	book := CreateBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
	}
	book := models.GetBook(ID)
	toBeUpdatedErr := utils.ParseBody(r, book)
	if toBeUpdatedErr != nil {
		fmt.Println(toBeUpdatedErr)
	}
	updatedBook := book.UpdateBook()
	res, _ := json.Marshal(updatedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
	}
	book := models.GetBook(ID)
	book.DeleteBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
