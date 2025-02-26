package routes

// golang has an absolute path so dont use ./ or / ....
import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/noogler-eng/go-lang/mysql/pkg/controllers"
)


func BookRouter() {
    r := mux.NewRouter()
    r.HandleFunc("/", CreateBooks);
    r.HandleFunc("/products", ProductsHandler)
    r.HandleFunc("/articles", ArticlesHandler)
}