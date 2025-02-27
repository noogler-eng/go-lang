package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/models"
)


func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	books := models.GetAllBooks();
	if books == nil {
		http.Error(w, `{"error": "No books found"}`, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(books);
}
