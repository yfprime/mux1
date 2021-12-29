package main
import(
	"context"
	"fmt"
	"net/http"
	"math/rand"
	"strconv"
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
//To get all books
func getBooks(w http.ResponseWriter, r *http.Request){


}
//To get a single book
func getBooks()(w http.ResponseWriter, r *http.Request){
	 

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
	// Initialize the router
	router := mux.NewRouter()

	//Handling Router: Establishing Endpoints

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", r)
	log.Fatal(http.ListenAndServe("8000",r))
}