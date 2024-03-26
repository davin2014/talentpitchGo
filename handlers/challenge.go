// El paquete handlers contiene los controladores de las rutas HTTP.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"errors"
	

	"talentpitchGo/models"     // modelos de datos
	"talentpitchGo/repository" // operaciones de base de datos
	"talentpitchGo/server"     // configuración del servidor

	
	"github.com/segmentio/ksuid" // para generar IDs únicos
	"github.com/gorilla/mux" // enrutador HTTP
)


type ChallengeRequest  struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  int    `json:"difficulty"`
	UserID      string `json:"user_id"`
}

type ChallengeResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
}


// crear es un controlador que maneja la creación de un nuevo usuario.
func CreateChallengeHandler(s server.Server) http.HandlerFunc {
		// Retornar la función del controlador
		return func(w http.ResponseWriter, r *http.Request) {
			// Decodificar el cuerpo de la solicitud
			var request ChallengeRequest
			
			// Decodificar el cuerpo de la solicitud
			err := json.NewDecoder(r.Body).Decode(&request)
			// Verificar si hubo un error decodificando el cuerpo de la solicitud
			if err != nil {
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
					// Si es otro error, maneja ese error
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
			
			// Generar un nuevo ID único
			id, err := ksuid.NewRandom()	
			if err != nil {
				// Retornar un error interno del servidor
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Crear una nueva estructura de usuario
			var challenge = models.Challenge{
				Id:       id.String(),
				Title:    request.Title,
				Description: request.Description,
				Difficulty: request.Difficulty,
				UserID: request.UserID,
			}
	
			
			// Insertar el usuario en la base de datos
			err = repository.InsertChallenge(r.Context(), &challenge)
			// Verificar si hubo un error insertando el usuario
			if err != nil {
				// Retornar un error interno del servidor
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Retornar la respuesta
			w.Header().Set("Content-Type", "application/json")
			// Codificar la respuesta
			json.NewEncoder(w).Encode(ChallengeResponse{
				Id: challenge.Id, 
				Title: challenge.Title,
			})
	}
}


// update es un controlador que maneja la actualización de un usuario existente.
func UpdateChallengeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud
		var request ChallengeRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		// Verificar si hubo un error decodificando el cuerpo de la solicitud
		if err != nil {
			// Retornar un error de solicitud incorrecta
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Obtener el ID del usuario de los parámetros de la ruta
		id := mux.Vars(r)["id"]
		// Crear una nueva estructura de usuario
		var challenge = models.Challenge{
			Id:       id,
			Title:    request.Title,
			Description: request.Description,
			Difficulty: request.Difficulty,
			UserID: request.UserID,
		}
		// Actualizar el usuario en la base de datos
		err = repository.UpdateChallenge(r.Context(), id, &challenge)
		// Verificar si hubo un error actualizando el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(ChallengeResponse{
			Id: challenge.Id, 
			Title: challenge.Title,
		})
	}
}


// delete es un controlador que maneja la eliminación de un usuario existente.
func DeleteChallengeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del usuario de los parámetros de la ruta
		id := mux.Vars(r)["id"]
		// Verificar si el usuario existe en la base de datos
		_, err := repository.GetChallengeById(r.Context(), id)
		// Verificar si hubo un error obteniendo el usuario
		if err != nil {
			if errors.Is(err, fmt.Errorf("no challenge found with id %s", id)) {
				http.Error(w, "Reto no encontrado", http.StatusNotFound)
			} else {
				// Si es otro error, maneja ese error
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Eliminar el usuario de la base de datos
		err = repository.DeleteChallenge(r.Context(), id)
		// Verificar si hubo un error eliminando el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		

		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(map[string]string{"message": "Challenge deleted"})
	}
}


// list es un controlador que maneja la obtención de una lista de usuarios.
func ListChallengesHandler(s server.Server) http.HandlerFunc {
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
		// Obtener la lista de usuarios de la base de datos
		challenges, total, err := repository.GetChallenges(r.Context(), pageNum, pageSizeNum)
		// Verificar si hubo un error obteniendo la lista de usuarios
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Crear una estructura de respuesta
		response := map[string]interface{} {
			"challenges": challenges,
			"total": total,
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(response)
	}
}


// get es un controlador que maneja la obtención de un usuario existente.
func GetChallengeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del usuario de los parámetros de la ruta
		id := mux.Vars(r)["id"]
		// Obtener el usuario de la base de datos
		challenge, err := repository.GetChallengeById(r.Context(), id)
		// Verificar si hubo un error obteniendo el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(challenge)
	}
}