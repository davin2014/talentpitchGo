// El paquete handlers contiene los controladores de las rutas HTTP.
package handlers

import (
	"encoding/json"
	"log"
    "strconv"
	"net/http"

	"errors"
	"fmt"

	"talentpitchGo/models"     // modelos de datos
	"talentpitchGo/repository" // operaciones de base de datos
	"talentpitchGo/server"     // configuración del servidor
    "github.com/gorilla/mux" // enrutador HTTP
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

// GetCompanyHandler es un controlador que maneja la obtención de una empresa por su ID.
func GetCompanyHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID de la empresa de los parámetros de la URL
		id := r.URL.Query().Get("id")
		// Verificar si el ID de la empresa está vacío
		if id == "" {
			// Retornar un error de solicitud incorrecta
			http.Error(w, "ID de empresa vacío", http.StatusBadRequest)
			return
		}

		// Obtener la empresa por su ID
		company, err := repository.GetCompanyById(r.Context(), id)
		// Verificar si hubo un error obteniendo la empresa
		if err != nil {
			// Si hay un error, verifica si es porque la empresa no fue encontrada
			if errors.Is(err, fmt.Errorf("no company found with id %s", id)) {
				http.Error(w, "Empresa no encontrada", http.StatusNotFound)
			} else {
				// Si hay un error diferente, retornar un error interno del servidor
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
			return
		}

		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(company)
	}
}

// UpdateCompanyHandler es un controlador que maneja la actualización de una empresa por su ID.
func UpdateCompanyHandler(s server.Server) http.HandlerFunc {
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
		// Obtener el ID del usuario de la URL
		id := mux.Vars(r)["id"]

		// Obtener la empresa por su ID
		company, err := repository.GetCompanyById(r.Context(), id)
		// Verificar si hubo un error obteniendo la empresa
		if err != nil {
			// Si hay un error, verifica si es porque la empresa no fue encontrada
			if errors.Is(err, fmt.Errorf("no company found with id %s", request.Id)) {
				log.Printf("Company not found: %v", err)
				http.Error(w, "Empresa no encontrada", http.StatusNotFound)
			} else {
				log.Printf("Error getting company: %v", err)
				// Si hay un error diferente, retornar un error interno del servidor
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
			return
		}

		

		// Crear una nueva estructura de empresa con los datos actualizados
		companyReq := models.Company{
			Id:        request.Id,
			Name:      request.Name,
			ImagePath: request.ImagePath,
			Location:  request.Location,
			Industry:  request.Industry,
			UserID:    request.UserID,
		}

		
		// Guardar la empresa actualizada en la base de datos
		err = repository.UpdateCompany(r.Context(), id, &companyReq)
		// Verificar si hubo un error guardando la empresa en la base de datos
		if err != nil {
			// Loggear el error
			log.Printf("Error updating company: %v", err)
			// Retornar un error interno del servidor
			http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			return
		}

		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(company)
	}
}

// DeleteCompanyHandler es un controlador que maneja la eliminación de una empresa por su ID.
func DeleteCompanyHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID de la empresa de los parámetros de la URL
		id := mux.Vars(r)["id"]
		// Verificar si el ID de la empresa está vacío
		if id == "" {
			// Retornar un error de solicitud incorrecta
			http.Error(w, "ID de empresa vacío", http.StatusBadRequest)
			return
		}

		// Eliminar la empresa por su ID
		err := repository.DeleteCompany(r.Context(), id)
		// Verificar si hubo un error eliminando la empresa
		if err != nil {
			// Si hay un error, verifica si es porque la empresa no fue encontrada
			if errors.Is(err, fmt.Errorf("no company found with id %s", id)) {
				http.Error(w, "Empresa no encontrada", http.StatusNotFound)
			} else {
				// Si hay un error diferente, retornar un error interno del servidor
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
			return
		}

		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode("Empresa eliminada exitosamente")
	}
}


// ListCompaniesHandler es un controlador que maneja la obtención de todas las empresas.
func ListCompaniesHandler(s server.Server) http.HandlerFunc {
// Retornar la función del controlador
return func(w http.ResponseWriter, r *http.Request) {
	// Obtener la página de la URL
	page := r.URL.Query().Get("page")
	// Convertir page a int
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		log.Print("Error converting page to int")
		// Manejar el error
	}
	// Obtener el tamaño de la página de la URL
	pageSize := r.URL.Query().Get("pageSize")
	// Convertir pageSize a int
	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		log.Print("Error converting pageSize to int")
		// Manejar el error
	}
	// Obtener la lista de usuarios
	companies, total, err := repository.GetCompanies(r.Context(), pageNum, pageSizeNum )
	// Verificar si hubo un error obteniendo la lista de usuarios
	if err != nil {
		// Retornar un error interno del servidor
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Crear un mapa con la lista de usuarios y el total
	response := map[string]interface{}{
		"companies": companies,
		"total": total,
	}
	// Retornar la respuesta
	w.Header().Set("Content-Type", "application/json")
	// Codificar la respuesta
	json.NewEncoder(w).Encode(response)
}
}