// El paquete server proporciona la configuración y el inicio del servidor.
package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"talentpitchGo/database" 
	"talentpitchGo/repository"
	"github.com/gorilla/mux"
)

// Config es la estructura de configuración del servidor.
type Config struct {
	Port string // Puerto del servidor
	JWTSecret string // Secreto para firmar el token JWT
	DatabaseURL string // URL de la base de datos
}

// Server es una interfaz que define las operaciones del servidor.
type Server interface {
	Config() *Config // Obtener la configuración del servidor
}

// Broker es una estructura que contiene la configuración y el enrutador del servidor.
type Broker struct {
	config *Config // Configuración del servidor
	router  *mux.Router // Enrutador del servidor
}

// Config es una función que obtiene la configuración del servidor.
func (b *Broker) Config() *Config {
	return b.config
}


// NewServer es una función que crea un nuevo servidor con la configuración dada.
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	// Verificar si el puerto está vacío
	if config.Port == "" {
		// Retornar un error de puerto requerido
		return nil, errors.New("port is required")
	}
	// Verificar si el secreto está vacío
	if config.JWTSecret == "" {
		// Retornar un error de secreto requerido
		return nil, errors.New("secret is required")
	}
	// Verificar si la URL de la base de datos está vacía
	if config.DatabaseURL == "" {
		// Retornar un error de base de datos requerida
		return nil, errors.New("database is required")
	}
	// Crear un nuevo servidor
	broker := &Broker{
		config: config,
		router:  mux.NewRouter(),
	}
	// Retornar el servidor creado
	return broker, nil
}

// Start es una función que inicia el servidor y enlaza las rutas con el enrutador.
func (b *Broker) Start(binder func (s Server, r *mux.Router )) {
	// Enlazar las rutas con el enrutador
	b.router = mux.NewRouter()
	// Iniciar el servidor
	binder(b, b.router)
	// Crear un nuevo repositorio de base de datos
	repo, err := database.NewPostgresRepository(b.Config().DatabaseURL)
	// Verificar si hubo un error creando el repositorio
	if err != nil {
		// Loggear el error
		log.Fatal("NewRepository", err)
	}
	// Establecer el repositorio de usuario
	repository.SetUserRepository(repo)
	// Establecer el repositorio de reto
	repository.SetChallengeRepository(repo)
	// Establecer el repositorio de empresa
	repository.SetCompanyRepository(repo)
	// Loggear el inicio del servidor
	log.Printf("Server is running on port %s", b.Config().Port)
	// Iniciar el servidor
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		// Loggear el error
		log.Fatal("ListenAndServe",err)
	}
}	