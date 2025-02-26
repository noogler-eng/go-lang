package routes

// golang has an absolute path so dont use ./ or / ....
// directly copy from github as a absoulte path
import (
	"github.com/gorilla/mux"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/controllers"
)

// always export things comes in caps
var BookRouter = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET");
	router.HandleFunc("/books/${id}", controllers.GetBooksById).Methods("GET");
    router.HandleFunc("/books", controllers.CreateBooks).Methods("POST");
    router.HandleFunc("/books/${id}", controllers.UpdateBook).Methods("PUT");
    router.HandleFunc("/books/${id}", controllers.DeleteBook).Methods("DELETE");
}