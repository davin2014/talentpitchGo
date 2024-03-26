// Descripcion: En este archivo se define la estructura de los usuarios
package models

// User es la estructura de los datos de un usuario
type User struct {
	Id       string  `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}