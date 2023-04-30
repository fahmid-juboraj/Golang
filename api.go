package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// Define a struct to represent a book
type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Define an in-memory slice to store books
var books []Book

// Define a handler function to handle requests to the /books endpoint
func booksHandler(w http.ResponseWriter, r *http.Request) {
    // Handle GET requests
    if r.Method == "GET" {
        // Marshal the books slice to JSON
        booksJson, err := json.Marshal(books)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // Set the response content-type header to application/json
        w.Header().Set("Content-Type", "application/json")
        // Write the JSON response to the http.ResponseWriter
        w.Write(booksJson)
    }
    // Handle POST requests
    if r.Method == "POST" {
        // Parse the request body into a Book struct
        var book Book
        err := json.NewDecoder(r.Body).Decode(&book)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        // Add the new book to the books slice
        books = append(books, book)
        // Marshal the new book to JSON
        bookJson, err := json.Marshal(book)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // Set the response content-type header to application/json
        w.Header().Set("Content-Type", "application/json")
        // Write the JSON response to the http.ResponseWriter
        w.Write(bookJson)
    }
}

func main() {
    // Add some initial books to the books slice
    books = []Book{
        Book{ID: 1, Title: "The Catcher in the Rye", Author: "J.D. Salinger"},
        Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
    }

    // Define a new http.ServeMux
    mux := http.NewServeMux()

    // Map the /books endpoint to the booksHandler function
    mux.HandleFunc("/books", booksHandler)

    // Start the server on port 8080
    log.Println("Listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
