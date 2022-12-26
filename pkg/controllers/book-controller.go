package controllers

import (
	"book-management-system/pkg/models"
	"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iamtonmoy0/book-management-system/pkg/models"
	"github.com/iamtonmoy0/book-management-system/pkg/utils"
)

var NewBook models.Book
func getBook(w http.ResponseWriter , r *http.Request){
	newBooks:= models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter ,r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID ,err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, := models.GetBookById(ID)
	res, :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter , r *http.Request){

	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:=CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter ,r * http.Request){
	vars := mux.Vars(r)
	bookId := Vars["bookId"]
	ID ,err := strconv.ParseInt(bookId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")

	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type" ,"pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func UpdateBook(w http.ResponseWriter ,r * http.Request){


	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars["bookId"]
	Id , err := strconv.ParseInt(bookId ,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails , bd :=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
}


