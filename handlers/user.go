// El paquete handlers contiene los controladores de las rutas HTTP.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"fmt"
	"errors"

	"talentpitchGo/models"     // modelos de datos
	"talentpitchGo/repository" // operaciones de base de datos
	"talentpitchGo/server"     // configuración del servidor

	"github.com/golang-jwt/jwt"  // para manejar JWT
	"github.com/segmentio/ksuid" // para generar IDs únicos
	"golang.org/x/crypto/bcrypt" // para hashear contraseñas
	"github.com/gorilla/mux" // enrutador HTTP
)

// HASH_COST es la complejidad del algoritmo de hash bcrypt.
const (
		HASH_COST = 8
)

// SingUpRequest es la estructura de los datos necesarios para registrar un nuevo usuario.
type SingUpRequest struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

// LoginResponse es la estructura de la respuesta del login
type LoginResponse struct {
	Token string `json:"token"`
}

// SingUpLoginRequest es la estructura de los datos necesarios para iniciar sesión.
type SingUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SingUpResponse es la estructura de la respuesta del registro
type SingUpResponse struct {
	Id 	 string  `json:"id"`
	Email    string `json:"email"`
}

// SignUpHandler es el controlador para la ruta de registro
func SignUpHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud
		var request SingUpRequest
		
		// Decodificar el cuerpo de la solicitud
		err := json.NewDecoder(r.Body).Decode(&request)
		// Verificar si hubo un error decodificando el cuerpo de la solicitud
		if err != nil {
			// Retornar un error de solicitud incorrecta
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		

		// Hashear la contraseña
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)
		// Verificar si hubo un error hasheando la contraseña
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		var user = models.User{
			Id:       id.String(),
			Fullname: request.Fullname,
			Email:    request.Email,
			Password: string(hashedPassword),
		}

		
		// Insertar el usuario en la base de datos
		err = repository.InsertUser(r.Context(), &user)
		// Verificar si hubo un error insertando el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(SingUpResponse{
			Id: user.Id, 
			Email: user.Email,
		})
}
}

// UpdateUserHandler es el controlador para la ruta de actualización de usuario
func UpdateUserHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud
		var request UpdateUserRequest
		// Decodificar el cuerpo de la solicitud
		err := json.NewDecoder(r.Body).Decode(&request)
		// Verificar si hubo un error decodificando el cuerpo de la solicitud
		if err != nil {
			// Retornar un error de solicitud incorrecta
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Obtener el ID del usuario de la URL
		id := mux.Vars(r)["id"]
		
		// Crear una nueva estructura de usuario
		var user = models.User{
			Id:       request.Id,
			Fullname: request.Fullname,
			Email:    request.Email,
		}
		// Insertar el usuario en la base de datos
	    err = repository.UpdateUser(r.Context(), id, &user)
		// Verificar si hubo un error insertando el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(SingUpResponse{
			Id: user.Id,
			Email: user.Email,
		})
	}
}

// DeleteUserHandler es el controlador para la ruta de eliminación de usuario
func DeleteUserHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del usuario de la URL
		id := mux.Vars(r)["id"]

		_, err := repository.GetUserById(r.Context(), id)
		// Verificar si hubo un error obteniendo el usuario
		// Verificar si hubo un error obteniendo el usuario
		if err != nil {
			// Si hay un error, verifica si es porque el usuario no fue encontrado
			if errors.Is(err, fmt.Errorf("no user found with id %s", id)) {
				http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			} else {
				// Si es otro error, maneja ese error
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		// Eliminar el usuario de la base de datos
		err = repository.DeleteUser(r.Context(), id)
		// Verificar si hubo un error eliminando el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
	}
}

// GetUserHandler es el controlador para la ruta de usuario
func GetUsersHandler(s server.Server) http.HandlerFunc {
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
		users, total, err := repository.GetUsers(r.Context(), pageNum, pageSizeNum )
		// Verificar si hubo un error obteniendo la lista de usuarios
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Crear un mapa con la lista de usuarios y el total
		response := map[string]interface{}{
			"users": users,
			"total": total,
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(response)
	}
}


// LoginHandler es el controlador para la ruta de inicio de sesión
func LoginHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud
		var request SingUpLoginRequest
		// Decodificar el cuerpo de la solicitud
		err := json.NewDecoder(r.Body).Decode(&request)
		// Verificar si hubo un error decodificando el cuerpo de la solicitud
		if err != nil {
			// Retornar un error de solicitud incorrecta
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Obtener el usuario por email
		user, err := repository.GetUserByEmail(r.Context(), request.Email)
		// Verificar si hubo un error obteniendo el usuario
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Verificar si el usuario es nulo
		if user == nil {
			// Retornar un error de credenciales inválidas
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		// Comparar la contraseña hasheada con la contraseña proporcionada
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		// Verificar si hubo un error comparando las contraseñas
		if  err != nil {
			// Retornar un error de credenciales inválidas
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		// Crear una nueva estructura de claims
		claims := models.AppClaims{
			UserId: user.Id,
            StandardClaims : jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		}
		// Crear un nuevo token con los claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Firmar el token con el secreto JWT
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		// Verificar si hubo un error firmando el token
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(LoginResponse{
			Token: tokenString,
		})
	}
}

// MeHandler es el controlador para la ruta de usuario
func MeHandler(s server.Server) http.HandlerFunc {
	// Retornar la función del controlador
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el token de la cabecera de autorización
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		// Parsear el token con los claims
        token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Retornar el secreto JWT de la configuración del servidor
			return []byte(s.Config().JWTSecret), nil
		})
		// Verificar si hubo un error parseando el token
		if err != nil {
			// Retornar un error de no autorizado
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Obtener los claims del token
		claims, ok := token.Claims.(*models.AppClaims)
		// Verificar si los claims son válidos
		if !ok && !token.Valid {
			// Retornar un error de token inválido
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// Obtener el usuario por su ID
		user, err := repository.GetUserById(r.Context(), claims.UserId)
		if err != nil {
			// Retornar un error interno del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Retornar la respuesta
		w.Header().Set("Content-Type", "application/json")
		// Codificar la respuesta
		json.NewEncoder(w).Encode(user)
	}
}