package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/utils"
)

// finding the book with that id
// remove the books
func DeleteBook(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
	id := vars["id"];

	// logic for deletion of the book form the database

	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(utils.Msg{Msg: "book has been deleted!"})
}
