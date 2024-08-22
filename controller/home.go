package controller

import (
	"libro-electronico/helper"
	"libro-electronico/helper/at"
	"libro-electronico/model"
	"net/http"
)

func GetHome(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Response = at.GetIPaddress()
	helper.WriteJSON(respw, http.StatusOK, resp)
}


func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Response = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
}