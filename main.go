package main
import(
	"context"
	"fmt"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	
)

func main(){ 
	// Initialize the router
	router := mux.NewRouter()

	//Handling Router: Establishing Endpoints

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	router.HandleFunc("/api/books", createBooks).Methods("POST")
	router.HandleFunc("/api/books{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/books{id}", deleteBook).Methods("DELETE")
}