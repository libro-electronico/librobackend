package routes

import (
	"libro-electronico/config"
	"libro-electronico/controller"
	"net/http"
)

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() {
	// Book routes
	http.HandleFunc("/api/books", withCORS(controller.GetBooks))
	http.HandleFunc("/api/books/create", withCORS(controller.CreateBook))
	http.HandleFunc("/api/books/update", withCORS(controller.UpdateBook))
	http.HandleFunc("/api/books/delete", withCORS(controller.DeleteBook))

	// User routes
	http.HandleFunc("/post/register", withCORS(controller.Register))
	http.HandleFunc("/post/login", withCORS(controller.Login))
}
