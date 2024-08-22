package route

import (
	"net/http"

	"libro-electronico/config"

	"libro-electronico/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	// Set Access Control Headers
	if config.SetAccessControlHeaders(w, r) {
		return
	}

	config.SetEnv()

	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			controller.GetHome(w, r)
		case "/api/get/books":
			controller.GetBooks(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/api/post/books":
			controller.CreateBook(w, r)
		case "/post/register":
			controller.Register(w, r)
		case "/post/login":
			controller.Login(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPut:
		if r.URL.Path == "/put/books" {
			controller.UpdateBook(w, r)
		} else {
			controller.NotFound(w, r)
		}
	case http.MethodDelete:
		if r.URL.Path == "/delete/books" {
			controller.DeleteBook(w, r)
		} else {
			controller.NotFound(w, r)
		}
	default:
		controller.NotFound(w, r)
	}
}