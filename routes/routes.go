package routes

import (
	"libro-electronico/controller"
	"net/http"
)

func SetupRoutes() {
	// Book routes
	http.HandleFunc("/api/books", controller.GetBooks)
	http.HandleFunc("/api/books/create", controller.CreateBook)
	http.HandleFunc("/api/books/update", controller.UpdateBook)
	http.HandleFunc("/api/books/delete", controller.DeleteBook)

	// User routes
	http.HandleFunc("/post/register", controller.Register)
	http.HandleFunc("/post/login", controller.Login)
}