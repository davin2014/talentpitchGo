package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"talentpitchGo/handlers"
	"talentpitchGo/server"
	"testing"

)

// Config es la estructura de configuración del servidor.
type Config struct {
	Port string // Puerto del servidor
	JWTSecret string // Secreto para firmar el token JWT
	DatabaseURL string // URL de la base de datos
}

// MockServer is a mock implementation of the Server interface for testing.
type MockServer struct{}

func (s *MockServer) Config() *server.Config {
    return &server.Config{
        Port:        "8080",
        JWTSecret:   "secret",
        DatabaseURL: "database_url",
    }
}

// SingUpResponse es la estructura de la respuesta del registro
type SingUpResponse struct {
	Id 	 string  `json:"id"`
	Email    string `json:"email"`
}




func TestSignUpHandler(t *testing.T) {
    req, err := http.NewRequest("POST", "/signup", strings.NewReader(`{"fullname":"Test User","email":"test@example.com","password":"password"}`))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SignUpHandler(&MockServer{}))

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var response SingUpResponse
    json.Unmarshal(rr.Body.Bytes(), &response)

    if response.Email != "test@example.com" {
        t.Errorf("handler returned unexpected body: got %v want %v", response.Email, "test@example.com")
    }
}