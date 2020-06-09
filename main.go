package main

import(
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book information 
type Book struct {
	ID string `json:"Id","omitemptly"`
	Title string `json:"Title","omitemptly"`
	Author *Author `json: "Author","omitemptly"`
	} 
	
// Author Information
type Author struct {
	Firstname string `json:"Firstname","omitemply"`
	Lastname string  `json:"Lastname","omitemptly"`
}
	
var books []Book
	
//GetBooks Get list of books
func GetBooks(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(books)
	}
	
	//GetBook get just one book of list
func GetBook(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		for _, item := range books {
			if item.ID == vars["Id"] {
				json.NewEncoder(w).Encode(item)
				return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
	
// CreateBook creates a new book 
func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&books)
	book.ID = vars["Id"]
	books = append(books,book)
	json.NewEncoder(w).Encode(books)
}

	//DeleteBook Delete a book of list
func DeleteBook(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		for index,item := range books {
			if item.ID == vars["Id"] {
				books = append(books[:index], books[index+1:]...)
				break
			}
			json.NewEncoder(w).Encode(books)
		}
	}
	
	
func main() {
		router := mux.NewRouter()
		
		books = append(books, Book{ID:"1",Title:"O perigo de uma história única",
		Author: &Author{Firstname:"Chimamanda",Lastname:"Ngozi Adichie"}})

		books = append(books, Book{ID:"2",Title:"Becoming", 
		Author: &Author{Firstname:"Michelle",Lastname:"Obama"}})
		
		books = append(books, Book{ID:"3",Title:"Memórias da plantação: Episódios de racismo cotidiano", 
		Author: &Author{Firstname:"Grada",Lastname:"Kilomba"}})
		
		books = append(books, Book{ID:"4",Title:"Women, Race, & Class",
		Author: &Author{Firstname:"Angela",Lastname:"Davis"}})


		router.HandleFunc("/books", GetBooks).Methods("GET")
		router.HandleFunc("/books/{Id}", GetBook).Methods("GET")
		router.HandleFunc("/books/{Id}", CreateBook).Methods("POST")
		router.HandleFunc("/books/{Id}",DeleteBook).Methods("DELETE")

		log.Fatal(http.ListenAndServe(":8000", router))
}