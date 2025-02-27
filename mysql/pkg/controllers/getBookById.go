package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/models"
)

func GetBooksById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json");
	id, exists := mux.Vars(r)["Id"];
	if !exists {
		http.Error(w, `{"error": "ID parameter is missing"}`, http.StatusBadRequest)
		return
	}
	idInInt, err := strconv.Atoi(id);
	if err != nil {
		http.Error(w, `{"error": "Invalid ID format"}`, http.StatusBadRequest)
		return
	}

	book, _ := models.GetBookById(int64(idInInt));
	json.NewEncoder(w).Encode(book)
}
