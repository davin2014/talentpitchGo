// El paquete main es el punto de entrada de la aplicación.
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"talentpitchGo/handlers" // controladores de rutas HTTP
	"talentpitchGo/middleware" // middleware de autenticación
	"talentpitchGo/server" // configuración del servidor

	"github.com/gorilla/mux" // enrutador HTTP
	"github.com/joho/godotenv" // para cargar variables de entorno desde un archivo .env
)

// main es la función principal que inicia el servidor.
func main() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load(".env")

	// Verificar si hubo un error cargando el archivo .env
	if err != nil {
		// Imprimir un mensaje de error y salir
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener el puerto, el secreto JWT y la URL de la base de datos desde las variables de entorno
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Crear un nuevo servidor con la configuración dada
	s, err := server.NewServer(context.Background(), &server.Config{
		Port: PORT,
		JWTSecret: JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
    // Manejar el error si existe
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}
    // Iniciar el servidor
	s.Start(BindRoutes)

	
}

// BindRoutes es una función que enlaza las rutas HTTP con los controladores.
func BindRoutes(s server.Server, r *mux.Router) {
	// Usar el middleware de autenticación
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/deleteUser/{id}", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/updateUser/{id}", handlers.UpdateUserHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/users", handlers.GetUsersHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.GetUsersHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
//************************************************************************************************************************
//************************************************************* CHALLENGE ************************************************
//************************************************************************************************************************
	r.HandleFunc("/challenges", handlers.CreateChallengeHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/updateChallenge/{id}", handlers.UpdateChallengeHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/deleteChallenge/{id}", handlers.DeleteChallengeHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/challenges", handlers.ListChallengesHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/challenges/{id}", handlers.GetChallengeHandler(s)).Methods(http.MethodGet)
//************************************************************************************************************************
//************************************************************* COMPANY **************************************************
//************************************************************************************************************************
	r.HandleFunc("/companies", handlers.CreateCompanyHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/updateCompany/{id}", handlers.UpdateCompanyHandler(s)).Methods(http.MethodPut)
	//r.HandleFunc("/deleteCompany/{id}", handlers.DeleteCompanyHandler(s)).Methods(http.MethodDelete)
	//r.HandleFunc("/companies", handlers.ListCompaniesHandler(s)).Methods(http.MethodGet)
	//r.HandleFunc("/companies/{id}", handlers.GetCompanyHandler(s)).Methods(http.MethodGet)


}