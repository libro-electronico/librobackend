package controller

import (
	"encoding/json"
	"net/http"

	"libro-electronico/config"
	"libro-electronico/helper"
	"libro-electronico/helper/chicken"
	"libro-electronico/model"

	"gopkg.in/mgo.v2/bson"
)

// GetBooks retrieves all books from the database
func GetBooks(respw http.ResponseWriter, req *http.Request) {
    books, _ := chicken.GetAllDoc[[]model.Book](config.Mongoconn, "librobackend", bson.M{})
    helper.WriteJSON(respw, http.StatusOK, books)
}

// CreateBook handles the creation of a new book
func CreateBook(respw http.ResponseWriter, req *http.Request) {
    var book model.Book
    err := json.NewDecoder(req.Body).Decode(&book)
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    _, err = chicken.InsertOneDoc(config.Mongoconn, "librobackend", book)
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    books, err := chicken.GetAllDoc[[]model.Book](config.Mongoconn, "librobackend", bson.M{})
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    helper.WriteJSON(respw, http.StatusOK, books)
}

// UpdateBook handles the updating of a book's information
func UpdateBook(respw http.ResponseWriter, req *http.Request) {
    var book model.Book
    err := json.NewDecoder(req.Body).Decode(&book)
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    dt, err := chicken.GetOneDoc[model.Book](config.Mongoconn, "librobackend", bson.M{"_id": book.ID})
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    dt.Title = book.Title
    dt.Author = book.Author
    dt.Publisher = book.Publisher
    dt.PublishedAt = book.PublishedAt
    dt.ISBN = book.ISBN
    dt.Pages = book.Pages
    dt.Language = book.Language
    dt.Available = book.Available
    _, err = chicken.ReplaceOneDoc(config.Mongoconn, "librobackend", bson.M{"_id": book.ID}, dt)
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    helper.WriteJSON(respw, http.StatusOK, dt)
}

// DeleteBook handles the deletion of a book
func DeleteBook(respw http.ResponseWriter, req *http.Request) {
    var book model.Book
    err := json.NewDecoder(req.Body).Decode(&book)
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    _, err = chicken.DeleteOneDoc(config.Mongoconn, "librobackend", bson.M{"_id": book.ID})
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    books, err := chicken.GetAllDoc[[]model.Book](config.Mongoconn, "librobackend", bson.M{})
    if err != nil {
        var respn model.Response
        respn.Response = err.Error()
        helper.WriteJSON(respw, http.StatusForbidden, respn)
        return
    }
    helper.WriteJSON(respw, http.StatusOK, books)
}