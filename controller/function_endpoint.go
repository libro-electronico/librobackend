package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"libro-electronico/config"
	"libro-electronico/helper"
	"libro-electronico/helper/at"
	"libro-electronico/helper/atdb"
	"libro-electronico/model"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// EnsureBookIDExists checks if a book ID exists and generates a unique ID if necessary.
func EnsureBookIDExists(book *model.Book) error {
	isUnique := false
	bookCount := 1
	for !isUnique {
		potentialID := fmt.Sprintf("%s-%03d", book.ISBN, bookCount)
		existingBook, err := helper.GetOneDoc[model.Book](config.Mongoconn, "backendlibro", bson.M{
			"id": potentialID,
		})
		if err != nil && err != mongo.ErrNoDocuments {
			return fmt.Errorf("failed to check for existing ID: %v", err)
		}
		if existingBook.ID == "" {
			book.ID = potentialID
			isUnique = true
		} else {
			bookCount++
		}
	}
	return nil
}

// GetBooks fetches books based on the provided filters.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")
	author := query.Get("author")

	filterBooks := bson.M{}
	if title != "" {
		filterBooks["title"] = title
	}
	if author != "" {
		filterBooks["author"] = author
	}

	findOptions := options.Find().SetLimit(20)

	var books []model.Book
	collection := config.Mongoconn.Collection("backendlibro")

	cursor, err := collection.Find(context.Background(), filterBooks, findOptions)
	if err != nil {
		log.Printf("Error fetching books: %v", err)
		http.Error(w, "Error fetching books from the database", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &books); err != nil {
		log.Printf("Error decoding books: %v", err)
		http.Error(w, "Error decoding books", http.StatusInternalServerError)
		return
	}

	if len(books) == 0 {
		http.Error(w, "No books found matching the filters", http.StatusNotFound)
		return
	}

	at.WriteJSON(w, http.StatusOK, books)
}

// PostBook adds a new book to the database.
func PostBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if newBook.Title == "" || newBook.Author == "" {
		http.Error(w, "Title and Author cannot be empty", http.StatusBadRequest)
		return
	}

	// Ensure the book has a unique ID
	if err := EnsureBookIDExists(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert the new book into the database
	if _, err := atdb.InsertOneDoc(config.Mongoconn, "backendlibro", newBook); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	at.WriteJSON(w, http.StatusOK, newBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    var book model.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        at.WriteJSON(w, http.StatusBadRequest, err.Error())
        return
    }

    // Update the book based on its ID
    filter := bson.M{"id": book.ID}
    update := bson.M{
        "$set": bson.M{
            "title":        book.Title,
            "author":       book.Author,
            "publisher":    book.Publisher,
            "published_at": book.PublishedAt,
            "isbn":         book.ISBN,
            "pages":        book.Pages,
            "language":     book.Language,
            "available":    book.Available,
        },
    }

    // Update the book in the database
    result, err := atdb.UpdateDoc(config.Mongoconn, "backendlibro", filter, update)
    if err != nil {
        at.WriteJSON(w, http.StatusInternalServerError, err.Error())
        return
    }

    if result.MatchedCount == 0 {
        at.WriteJSON(w, http.StatusNotFound, "Book not found")
        return
    }

    at.WriteJSON(w, http.StatusOK, book)
}


// DeleteBook deletes a book from the database.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	filter := bson.M{"id": book.ID}

	// Check if the book exists
	existingBook, err := helper.GetOneDoc[model.Book](config.Mongoconn, "backendlibro", filter)
	if err != nil || existingBook.ID == "" {
		http.Error(w, fmt.Sprintf("No book found with ID: %s", book.ID), http.StatusNotFound)
		return
	}

	// Delete the book
	if err := atdb.DeleteOneDoc(config.Mongoconn, "backendlibro", filter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	at.WriteJSON(w, http.StatusOK, fmt.Sprintf("Book with ID: %s has been successfully deleted", book.ID))
}
