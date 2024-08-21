package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// connect db
type DBIngfo struct {
	DBString string
	DBName   string
}
//response error
type Response struct {
	Response string `json:"response"`
	Info     string `json:"info,omitempty"`
	Status   string `json:"status,omitempty"`
	Location string `json:"location,omitempty"`
}
// login and register for user
type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

// login request untuk user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// book type
type Book struct {
    ID          string  `json:"id"`          // ID unik untuk setiap buku
    Title       string  `json:"title"`       // Judul buku
    Author      string  `json:"author"`      // Penulis buku
    Publisher   string  `json:"publisher"`   // Penerbit buku
    PublishedAt string  `json:"published_at"`// Tanggal publikasi buku
    ISBN        string  `json:"isbn"`        // Nomor ISBN buku
    Pages       int     `json:"pages"`       // Jumlah halaman buku
    Language    string  `json:"language"`    // Bahasa buku
    Available   bool    `json:"available"`   // Status ketersediaan buku
}