package routes

import (
	"libro-electronico/controller"
	"net/http"
)

func SetupRoutes() {
	// Book routes
	http.HandleFunc("/get/books", controller.GetBooks)
	http.HandleFunc("/post/books/create", controller.CreateBook)
	http.HandleFunc("/put/books/update", controller.UpdateBook)
	http.HandleFunc("/delete/books/delete", controller.DeleteBook)

	// User routes
	http.HandleFunc("/post/register", controller.Register)
	http.HandleFunc("/post/login", controller.Login)
}