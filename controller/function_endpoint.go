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
	resp, _:= atdb.GetOneDoc[[]model.Book](config.Mongoconn, "backendlibro", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, resp)
	
}

func CreateBook(respw http.ResponseWriter, req *http.Request) {
	var item model.Book
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	_, err = atdb.InsertOneDoc(config.Mongoconn,"backendlibro",item)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	books, err  := atdb.GetAllDoc[[]model.Book](config.Mongoconn,"backendlibro",bson.M{})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	helper.WriteJSON(respw, http.StatusOK, books)
	
}

func UpdateBook(respw http.ResponseWriter, req *http.Request) {
	var item model.Book
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	dt, err:= atdb.GetOneDoc[model.Book](config.Mongoconn,"backendlibro",bson.M{"_id":item.ID})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	dt.JamOperasional = rute.JamOperasional
	dt.Rute = rute.Rute
	dt.Tarif = rute.Tarif
	_, err= atdb.ReplaceOneDoc(config.Mongoconn,"backendlibro",bson.M{"_id":item.ID},dt)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	helper.WriteJSON(respw, http.StatusOK, dt)
	
}

func DeleteBook(respw http.ResponseWriter, req *http.Request) {
	var item model.Book
	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	err = atdb.DeleteOneDoc(config.Mongoconn,"backendlibro",bson.M{"_id":item.ID})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	books, err  := atdb.GetAllDoc[[]model.Book](config.Mongoconn,"backendlibro",bson.M{"_id":item.ID})
	if err != nil {
		var respn model.Response
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
		
	}
	helper.WriteJSON(respw, http.StatusOK, books)

	
}
