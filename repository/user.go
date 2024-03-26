// El paquete repository proporciona las operaciones de base de datos para los usuarios.
package repository

import (
	"errors"
	"context"
	"talentpitchGo/models"
)

// UserRepository es una interfaz que define las operaciones de base de datos para los usuarios.
type UserRepository interface {
	// InsertUser es una función que inserta un nuevo usuario en la base de datos.
	InsertUser(ctx context.Context, user *models.User) error
	// UpdateUser es una función que actualiza un usuario en la base de datos.
	UpdateUser(ctx context.Context, id string, user *models.User) error
	// DeleteUser es una función que elimina un usuario de la base de datos por su ID.
	DeleteUser(ctx context.Context, id string) error
	// GetUsers es una función que obtiene una lista de usuarios de la base de datos.
	GetUsers(ctx context.Context, page int, pageSize int) ([]*models.User, int, error)
	// GetUserById es una función que obtiene un usuario de la base de datos por su ID.
	GetUserById(ctx context.Context, id string) (*models.User, error)
	// GetUserByEmail es una función que obtiene un usuario de la base de datos por su correo electrónico.
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// Close es una función que cierra la conexión a la base de datos.
	Close() error
}

// implementation es una variable que contiene la implementación de la interfaz UserRepository.
var implementation UserRepository

// SetUserRepository es una función que establece la implementación de la interfaz UserRepository.
func SetUserRepository(repo UserRepository) {
	// Establecer la implementación de la interfaz UserRepository
	implementation = repo
}


// InsertUser es una función que inserta un nuevo usuario en la base de datos.
func InsertUser(ctx context.Context, user *models.User) error {

	// Verificar que implementation no sea nil
    if implementation == nil {
        return errors.New("implementation cannot be nil")
    }
	
	// Verificar que el contexto no sea nil
    if ctx == nil {
        return errors.New("context cannot be nil")
    }

    // Verificar que el usuario no sea nil
    if user == nil {
        return errors.New("user cannot be nil")
    }

    // Verificar que los campos necesarios del usuario no estén vacíos
    if user.Email == "" {
        return errors.New("user email cannot be empty")
    }
    if user.Fullname == "" {
        return errors.New("user fullname cannot be empty")
    }
    if user.Password == "" {
        return errors.New("user password cannot be empty")
    }
	// Insertar un nuevo usuario en la base de datos
	return implementation.InsertUser(ctx, user)
}

// UpdateUser es una función que actualiza un usuario en la base de datos.
func UpdateUser(ctx context.Context ,id string, user *models.User) error {
	// Verificar que implementation no sea nil
	if implementation == nil {
		return  errors.New("implementation cannot be nil")
	}
	
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return  errors.New("context cannot be nil")
	}

	// Verificar que el usuario no sea nil
	if user == nil {
		return  errors.New("user cannot be nil")
	}

	// Verificar que los campos necesarios del usuario no estén vacíos
	if user.Email == "" {
		return  errors.New("user email cannot be empty")
	}
	if user.Fullname == "" {
		return errors.New("user fullname cannot be empty")
	}
	// Insertar un nuevo usuario en la base de datos
	return implementation.UpdateUser(ctx, id, user)
}

// DeleteUser es una función que elimina un usuario de la base de datos por su ID.
func DeleteUser(ctx context.Context, id string) error {
	// Verificar que implementation no sea nil
	if implementation == nil {
		return errors.New("implementation cannot be nil")
	}
	
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	// Eliminar un usuario de la base de datos por su ID
	return implementation.DeleteUser(ctx, id)
}


// GetUsers es una función que obtiene una lista de usuarios de la base de datos.
func GetUsers(ctx context.Context, page int, pageSize int) ([]*models.User, int, error) {
	// Obtener una lista de usuarios de la base de datos
	return implementation.GetUsers(ctx, page, pageSize)
}

// GetUserById es una función que obtiene un usuario de la base de datos por su ID.
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	// Obtener un usuario de la base de datos por su ID
	return implementation.GetUserById(ctx, id)
}


// GetUserByEmail es una función que obtiene un usuario de la base de datos por su correo electrónico.
func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	// Obtener un usuario de la base de datos por su correo electrónico
	return implementation.GetUserByEmail(ctx, email)
}


// Close es una función que cierra la conexión a la base de datos.
func Close() error {
	// Cerrar la conexión a la base de datos
	return implementation.Close()
}
