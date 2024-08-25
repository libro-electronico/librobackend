package controller

import (
	"context"
	"encoding/json"
	"io"
	"libro-electronico/config"
	"libro-electronico/model"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
        return
    }

    // Log the request body for debugging
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, `{"error": "Failed to read request body"}`, http.StatusInternalServerError)
        return
    }
    log.Println("Received request body:", string(body))

    // Decode the request body into the User struct
    var user model.User
    err = json.Unmarshal(body, &user)
    if err != nil {
        http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
        return
    }

    // Check for empty fields
    if user.Username == "" || user.Password == "" || user.Email == "" {
        http.Error(w, `{"error": "Missing required fields"}`, http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, `{"error": "Failed to hash password"}`, http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)
    user.ID = primitive.NewObjectID()

    // Insert user into MongoDB
    collection := config.Mongoconn.Collection("user_login")
    _, err = collection.InsertOne(context.Background(), user)
    if err != nil {
        http.Error(w, `{"error": "Failed to register user"}`, http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
        return
    }

    // Decode the request body into the LoginRequest struct
    var loginRequest model.LoginRequest
    err := json.NewDecoder(r.Body).Decode(&loginRequest)
    if err != nil {
        http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
        return
    }

    // Fetch user based on email and username to get the user ID
    collection := config.Mongoconn.Collection("user_login")
    var user model.User
    err = collection.FindOne(context.Background(), bson.M{"email": loginRequest.Email, "username": loginRequest.Username}).Decode(&user)
    if err != nil {
        http.Error(w, `{"error": "Invalid email or username"}`, http.StatusUnauthorized)
        return
    }

    // Compare the hashed password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
    if err != nil {
        http.Error(w, `{"error": "Invalid email or password"}`, http.StatusUnauthorized)
        return
    }

    response := map[string]interface{}{
        "message": "Login successful",
        "user": map[string]interface{}{
            "id":       user.ID.Hex(),
            "username": user.Username,
            "email":    user.Email,
        },
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}