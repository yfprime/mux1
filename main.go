package main

import (
	//"context"
	//"fmt"
	"log"
	"net/http"
	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

//Structure for Books, similar to an object in Object-Oriented Programming.

type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"` //*Author is it's own kind of structure, which we create below.
}

//Creating the Author Structure

type Author struct{ 
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	
}
//Init books variable as a slice Book structure. A slice is an array with variable length.
var books []Book
//To get all books
func getBooks(w http.ResponseWriter, r *http.Request){


}
//To get a single book
func getBook(w http.ResponseWriter, r *http.Request){
	 

}
//To create a new book
func createBook(w http.ResponseWriter, r *http.Request){
	 

}
//To update a book.
func updateBook(w http.ResponseWriter, r *http.Request){
	 

}
//To delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
	 

}

func main(){ 
	// Initialize the r
	r := mux.NewRouter()
	//Mock Data
	books = append(books, Book{ID:"1", Isbn:"74574",Title:"Book 1",Author: &Author{Firstname: "Yousuf", Lastname: "Farhan"}})
	//Handling r: Establishing Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", r)
	log.Fatal(http.ListenAndServe(":8000",r))
}