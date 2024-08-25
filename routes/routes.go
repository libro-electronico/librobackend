package route

import (
	"net/http"

	"libro-electronico/config"
	"libro-electronico/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	// Set Access Control Headers (CORS)
	if config.SetAccessControlHeaders(w, r) {
		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// Set environment variables if needed
	config.SetEnv()

	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPut:
		handlePut(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		controller.NotAllowed(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		controller.GetHome(w, r)
	case "/api/get/books":
		controller.GetBooks(w, r)
	default:
		controller.NotFound(w, r)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/post/books":
		controller.PostBook(w, r)
	case "/post/register":
		controller.Register(w, r)
	case "/post/login":
		controller.Login(w, r)
	default:
		controller.NotFound(w, r)
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/put/books":
		controller.UpdateBook(w, r)
	default:
		controller.NotFound(w, r)
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/delete/books":
		controller.DeleteBook(w, r)
	default:
		controller.NotFound(w, r)
	}
}
