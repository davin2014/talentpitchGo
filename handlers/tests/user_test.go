package server_test

import (
	"os"
	"log"
	"context"
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "talentpitchGo/handlers"
    "talentpitchGo/server"
	"talentpitchGo/database"
	"talentpitchGo/repository"

	"github.com/joho/godotenv" // para cargar variables de entorno desde un archivo .env
	"github.com/stretchr/testify/assert"
)

func TestSignUpHandler(t *testing.T) {

	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load("./../../.env_test")

	// Verificar si hubo un error cargando el archivo .env
	if err != nil {
		// Imprimir un mensaje de error y salir
		t.Fatalf("Error loading .env file: %v", err)
	}
	
	// Obtener el puerto, el secreto JWT y la URL de la base de datos desde las variables de entorno
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

    
    // Crear un servidor HTTP de prueba
    s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		JWTSecret: JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})

	// verificar si hubo un error creando el servidor HTTP de prueba
	if err != nil {
		// manejar el error
		t.Fatal(err)
	}

	
	// Crear un nuevo repositorio de base de datos
	repo, err := database.NewPostgresRepository(s.Config().DatabaseURL)
	// Verificar si hubo un error creando el repositorio
	if err != nil {
		// Loggear el error y fallar la prueba
		t.Fatalf("NewRepository: %v", err)
	}
	// Establecer el repositorio de usuario
	repository.SetUserRepository(repo)
    
    // Crear un controlador de registro
    handler := handlers.SignUpHandler(s)
	if handler == nil {
		t.Fatalf("expected handler to be not nil")
	}
    
    
	
    // Crear una solicitud de registro
    request := handlers.SingUpRequest{
        Fullname: "test",
        Email:    "test@example.com",
        Password: "password",
    }
    requestBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Error marshalling request: %v", err)
	}

	
    // Crear una solicitud HTTP de prueba
    req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
    
	
	// verificar si hubo un error creando la solicitud HTTP de prueba
	if err != nil {
		// manejar el error
		t.Fatal(err)
	}
    
	// Establecer el encabezado Content-Type en la solicitud
    req.Header.Set("Content-Type", "application/json")

    // Crear un ResponseWriter de prueba
    rr := httptest.NewRecorder()

	// verificar si hubo un error creando el ResponseWriter de prueba
	if rr == nil {
		// manejar el error
		t.Fatal(err)
	}

    
    // Llamar al controlador de registro
    handler.ServeHTTP(rr, req)

    // Comprobar que la respuesta tiene un código de estado 200
    assert.Equal(t, http.StatusOK, rr.Code)
    
    // Decodificar la respuesta
    var response handlers.SingUpResponse
    json.NewDecoder(rr.Body).Decode(&response)

    // Comprobar que la respuesta tiene un ID y un correo electrónico
    assert.NotEmpty(t, response.Id)
    assert.Equal(t, "test@example.com", response.Email)
}


func TestLoginHandler(t *testing.T) {
	
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load("./../../.env_test")

	// Verificar si hubo un error cargando el archivo .env
	if err != nil {
		// Imprimir un mensaje de error y salir
		t.Fatalf("Error loading .env file: %v", err)
	}
	
	// Obtener el puerto, el secreto JWT y la URL de la base de datos desde las variables de entorno
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	
	// Crear un servidor HTTP de prueba
	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		JWTSecret: JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})

	// verificar si hubo un error creando el servidor HTTP de prueba
	if err != nil {
		// manejar el error
		t.Fatal(err)
	}

	
	// Crear un nuevo repositorio de base de datos
	repo, err := database.NewPostgresRepository(s.Config().DatabaseURL)
	// Verificar si hubo un error creando el repositorio
	if err != nil {
		// Loggear el error y fallar la prueba
		t.Fatalf("NewRepository: %v", err)
	}
	// Establecer el repositorio de usuario
	repository.SetUserRepository(repo)
	
	// Crear un controlador de registro
	handler := handlers.LoginHandler(s)
	if handler == nil {
		t.Fatalf("expected handler to be not nil")
	}
	
	
	
	// Crear una solicitud de registro
	request := handlers.SingUpLoginRequest{
		Email:    "test@example.com",
        Password: "password",
	}
	requestBody, err := json.Marshal(request)

	// verificar si hubo un error creando la solicitud de registro
	if err != nil {
		// manejar el error
		t.Fatal(err)
	}

	// crear una solicitud HTTP de registro
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))

	// verificar si hubo un error creando la solicitud HTTP de registro
	if err != nil {
		// manejar el error
		t.Fatal(err)
	}

	// verificar si hubo un error creando la solicitud HTTP de prueba
	if err != nil {
		// manejar el error
		t.Fatal(err)
	}
    
	// Establecer el encabezado Content-Type en la solicitud
    req.Header.Set("Content-Type", "application/json")

    // Crear un ResponseWriter de prueba
    rr := httptest.NewRecorder()

	// verificar si hubo un error creando el ResponseWriter de prueba
	if rr == nil {
		// manejar el error
		t.Fatal(err)
	}

    
    // Llamar al controlador de registro
    handler.ServeHTTP(rr, req)

    // Comprobar que la respuesta tiene un código de estado 200
    assert.Equal(t, http.StatusOK, rr.Code)
}


