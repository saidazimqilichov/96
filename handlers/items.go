package handlers

import (
	"encoding/json"
	"https/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAllItems())
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	newItem := models.CreateItem(item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	item, found := models.GetItem(id)
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = id
	updatedItem := models.UpdateItem(item)
	json.NewEncoder(w).Encode(updatedItem)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	models.DeleteItem(id)
	w.WriteHeader(http.StatusNoContent)
}
