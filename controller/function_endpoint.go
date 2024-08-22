package controller

import (
	"encoding/json"
	"net/http"

	"libro-electronico/config"
	"libro-electronico/helper"
	"libro-electronico/helper/at"
	"libro-electronico/helper/atdb"
	"libro-electronico/model"

	"go.mongodb.org/mongo-driver/bson"
)

func GetBooks(respw http.ResponseWriter, req *http.Request) {
	books, err := atdb.GetAllDoc[[]model.Book](config.Mongoconn, "backendlibro", bson.M{})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, books)
}

func CreateBook(respw http.ResponseWriter, req *http.Request) {
	var item model.Book
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		var respn model.Response
		respn.Response = "Invalid request payload: " + err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	_, err = atdb.InsertOneDoc(config.Mongoconn, "backendlibro", item)
	if err != nil {
		var respn model.Response
		respn.Response = "Failed to insert book: " + err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}

	books, err := atdb.GetAllDoc[[]model.Book](config.Mongoconn, "backendlibro", bson.M{})
	if err != nil {
		var respn model.Response
		respn.Response = "Failed to fetch books: " + err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, books)
}

func UpdateBook(respw http.ResponseWriter, req *http.Request) {
	var item model.Book
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		var respn model.Response
		respn.Response = "Invalid request payload: " + err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	dt, err := atdb.GetOneDoc[model.Book](config.Mongoconn, "backendlibro", bson.M{"_id": item.ID})
	if err != nil {
		var respn model.Response
		respn.Response = "Book not found: " + err.Error()
		at.WriteJSON(respw, http.StatusNotFound, respn)
		return
	}

	dt.Title = item.Title
	dt.Author = item.Author
	dt.Publisher = item.Publisher
	dt.PublishedAt = item.PublishedAt
	dt.ISBN = item.ISBN
	dt.Pages = item.Pages
	dt.Language = item.Language
	dt.Available = item.Available

	_, err = atdb.ReplaceOneDoc(config.Mongoconn, "backendlibro", bson.M{"_id": item.ID}, dt)
	if err != nil {
		var respn model.Response
		respn.Response = "Failed to update book: " + err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, dt)
}

func DeleteBook(respw http.ResponseWriter, req *http.Request) {
	var item model.Book
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		var respn model.Response
		respn.Response = "Invalid request payload: " + err.Error()
		at.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}

	err = atdb.DeleteOneDoc(config.Mongoconn, "backendlibro", bson.M{"_id": item.ID})
	if err != nil {
		var respn model.Response
		respn.Response = "Failed to delete book: " + err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}

	books, err := atdb.GetAllDoc[[]model.Book](config.Mongoconn, "backendlibro", bson.M{})
	if err != nil {
		var respn model.Response
		respn.Response = "Failed to fetch books: " + err.Error()
		at.WriteJSON(respw, http.StatusInternalServerError, respn)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, books)
}
