package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hornlessmoose/crud-api/pkg/models"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var updatedBook models.Book

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	
	json.Unmarshal(body, &updatedBook)

	id, _ := strconv.Atoi(vars["id"])

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
		json.NewEncoder(w).Encode(http.StatusNotFound)
		return
	}

	if updatedBook.Title != "" {
		book.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Desc != "" {
		book.Desc = updatedBook.Desc
	}

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}