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
			controller.GetBooks(w, r) // Endpoint untuk GET request (Mendapatkan daftar buku)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/api/post/books":
			controller.PostBook(w, r) // Endpoint untuk POST request (Membuat buku baru)
		case "/post/register":
			controller.Register(w, r)
		case "/post/login":
			controller.Login(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPut:
		switch r.URL.Path {
		case "/api/put/books":
			controller.UpdateBook(w, r) // Endpoint untuk PUT request (Mengupdate buku)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodDelete:
		switch r.URL.Path {
		case "/api/delete/books":
			controller.DeleteBook(w, r) // Endpoint untuk DELETE request (Menghapus buku)
		default:
			controller.NotFound(w, r)
		}
	default:
		controller.NotFound(w, r)
	}
}
