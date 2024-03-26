// Descripcion : Archivo con la estructura de los claims del token jwt
package models

import "github.com/golang-jwt/jwt"

// AppClaims es la estructura de los claims del token JWT
type AppClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims 
}