package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/models"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/utils"
)

// finding the book with that id
// remove the books
func DeleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");

    vars := mux.Vars(r)
	id, exists := vars["id"];
	if !exists {
		http.Error(w, `{"error": "ID parameter is missing"}`, http.StatusBadRequest)
		return
	}

	idInInt, err := strconv.Atoi(id);
	if err != nil {
		http.Error(w, `{"error": "Invalid ID format"}`, http.StatusBadRequest)
		return
	}
	
	// logic for deletion of the book form the database
	models.DeleteBookById(int64(idInInt));
	json.NewEncoder(w).Encode(utils.Msg{Msg: "book has been deleted!"})
}
