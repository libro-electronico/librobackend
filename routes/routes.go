package routes

import (
	"libro-electronico/config"
	"libro-electronico/controller"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		controller.GetBooks(w, r)
	})
	http.HandleFunc("/api/books/create", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		controller.CreateBook(w, r)
	})
	http.HandleFunc("/api/books/update", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		controller.UpdateBook(w, r)
	})
	http.HandleFunc("/api/books/delete", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		controller.DeleteBook(w, r)
	})

	// User routes
	http.HandleFunc("/post/register", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		controller.Register(w, r)
	})
	http.HandleFunc("/post/login", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		controller.Login(w, r)
	})
}
