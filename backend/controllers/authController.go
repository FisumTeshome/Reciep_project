package controllers

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

// Credentials stores the username and password
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// saveUserToDatabase saves the username and hashed password to the database
func saveUserToDatabase(username, hashedPassword string) error {
    // Implement your database logic here
    // For now, we'll just return nil to simulate a successful save
    return nil
}

// Claims stores the JWT claims
type Claims struct {
    Username string `json:"username"`
    jwt.RegisteredClaims
}

// SignUp handles user registration
func SignUp(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // Save the username and hashed password to your database
    err = saveUserToDatabase(creds.Username, string(hashedPassword))
    if err != nil {
        http.Error(w, "Error saving user to database", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// SignIn handles user login
func SignIn(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Retrieve the hashed password from your database for the given username
    // Example: hashedPassword := getUserPassword(creds.Username)
    hashedPassword := "" // Replace with actual retrieval logic
    // ...

    // Compare the provided password with the hashed password
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Username: creds.Username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })
}

// Welcome handles authenticated requests
func Welcome(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    tokenStr := cookie.Value
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        if err == jwt.ErrSignatureInvalid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    if !token.Valid {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    w.Write([]byte("Welcome " + claims.Username))
}