package config

import (
	"net/http"
)

// Allowed Origins should match the domain from which requests are allowed
var Origins = []string{
	"https://libro-electronico.github.io",
}

func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "https://libro-electronico.github.io")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	// Set CORS headers for the main request
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "https://libro-electronico.github.io")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	return false
}
