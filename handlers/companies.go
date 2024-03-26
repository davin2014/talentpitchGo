// El paquete handlers contiene los controladores de las rutas HTTP.
package handlers

import (
	"encoding/json"
	"log"

	"net/http"

	"errors"
	"fmt"

	"talentpitchGo/models"     // modelos de datos
	"talentpitchGo/repository" // operaciones de base de datos
	"talentpitchGo/server"     // configuración del servidor

	"github.com/segmentio/ksuid" // para generar IDs únicos
)

// CompanyRequest es una estructura que representa la solicitud de creación de una empresa.
type CompanyRequest  struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ImagePath   string `json:"image_path"`
	Location    string `json:"location"`
	Industry    string `json:"industry"`
	UserID      string `json:"user_id"`
}

// CompanyResponse es una estructura que representa la respuesta de creación de una empresa.
type CompanyResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
}

// CreateCompanyHandler es un controlador que maneja la creación de una nueva empresa.
func CreateCompanyHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud
		var request CompanyRequest
		
		// Decodificar el cuerpo de la solicitud
		err := json.NewDecoder(r.Body).Decode(&request)
		// Verificar si hubo un error decodificando el cuerpo de la solicitud
		if err != nil {
			// Loggear el error
			log.Printf("Error decoding request: %v", err)
			// Retornar un error de solicitud incorrecta
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = repository.GetUserById(r.Context(), request.UserID)
		// Verificar si hubo un error obteniendo el usuario
		// Verificar si hubo un error obteniendo el usuario
		if err != nil {
			// Si hay un error, verifica si es porque el usuario no fue encontrado
			if errors.Is(err, fmt.Errorf("no user found with id %s", request.UserID)) {
				http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			} else {
				// Si hay un error diferente, retornar un error interno del servidor
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
			return
		}

		// Crear una nueva empresa
		company := models.Company{
			Id:        ksuid.New().String(),
			Name:      request.Name,
			ImagePath: request.ImagePath,
			Location:  request.Location,
			Industry:  request.Industry,
			UserID:    request.UserID,
		}

		// Guardar la empresa en la base de datos
		err = repository.InsertCompany(r.Context(), &company)
		// Verificar si hubo un error guardando la empresa en la base de datos
		if err != nil {
			// Loggear el error
			log.Printf("Error inserting company: %v", err)
			// Retornar un error interno del servidor
			http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			return
		}

		
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(CompanyResponse{
			Id: company.Id, 
			Name: company.Name,
		})
	}
}