package main

import (
	//"context"
	"sync"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	//"time"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"html/template"
)

//Structure for Books, similar to an object in Object-Oriented Programming.

type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"` //*Author is it's own kind of structure, which we create below.
}
var wg sync.WaitGroup 
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
//To get a single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through books and find id
	for _, item := range books { 
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
	json.NewEncoder(w).Encode(&Book{})
	defer fmt.Println("Here is a list of all the books in storage.")

}
//To create a new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID =  strconv.Itoa(rand.Intn(1000000)) //Mock ID - not unique
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

	 

}
//To update a book.
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{ 
		if item.ID == params["id"]{
		books = append(books[:index],books[index+1:]...)
		var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID =   params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return
		}

	}
	json.NewEncoder(w).Encode(books)	 

}
//To delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "application/json")
	 params := mux.Vars(r)
	 for index, item := range books{ 
		 if item.ID == params["id"]{
		 books = append(books[:index],books[index+1:]...)
		 break
		 }

	 }
	 json.NewEncoder(w).Encode(books)

}
var templates *template.Template
func main(){ 
	// Initialize the r
	r := mux.NewRouter()
	//Mock Data
	books = append(books, Book{ID:"2", Isbn:"7adsaddas4574",Title:"Book 2",Author: &Author{ID:"2",Firstname: "Andrew", Lastname: "Paul"}})

	books = append(books, Book{ID:"1", Isbn:"74574",Title:"Book 1",Author: &Author{ID:"1",Firstname: "Yousuf", Lastname: "Farhan"}})
	//Ensuring that only templaters from the templates folder are parsed and instantiated.
	templates =template.Must(template.ParseGlob("templates/*.html"))
	//Handling r: Establishing Endpoints
	r.HandleFunc("/", indexTemplate)
	http.Handle("/",r)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", r)
	log.Fatal(http.ListenAndServe(":8000",r))
}
func indexTemplate(w http.ResponseWriter, r* http.Request){ 
	templates.ExecuteTemplate(w, "index.html",nil)
}