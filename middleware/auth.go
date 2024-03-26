// Description: This file contains the middleware to check if the request has a valid JWT token
package middleware

import (
	"net/http"
	"strings"
	"talentpitchGo/server"
	"talentpitchGo/models"
	"github.com/golang-jwt/jwt"
)
	
// NO_AUTH_NEEDED esta variable contiene las rutas que no necesitan autenticación
var (
	NO_AUTH_NEEDED = []string{
		"/signup", 
		"/login",
	}
)

// shouldCheckToken es una función para verificar si el token debe ser verificado para la ruta dada
func shouldCheckToken(route string) bool {
	// Iterar sobre las rutas que no necesitan autenticación
	for _, p := range NO_AUTH_NEEDED {
		// Verificar si la ruta contiene la ruta actual
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

// CheckAuthMiddleware es un middleware para verificar si la solicitud tiene un token JWT válido
func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler{
	// Retornar un http.HandlerFunc
	return func(next http.Handler) http.Handler {
		// Retornar un http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verificar si el token no debe ser verificado
			if !shouldCheckToken(r.URL.Path) {
				// Llamar al siguiente manejador
				next.ServeHTTP(w, r)
				return
			}
			// Obtener el token de la cabecera de autorización
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			// Verificar si el token está vacío
			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Retornar el secreto JWT de la configuración del servidor
				return []byte(s.Config().JWTSecret), nil
			})
			// Verificar si hubo un error al verificar el token
			if err != nil {
				// Retornar un error de no autorizado
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			// Llamar al siguiente manejador
			next.ServeHTTP(w, r)
		})
	}
}