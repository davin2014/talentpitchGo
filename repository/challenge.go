// El paquete repository proporciona las operaciones de base de datos para los usuarios.
package repository

import (
	"errors"
	"context"
	"talentpitchGo/models"
)


// ChallengeRepository es una interfaz que define las operaciones de base de datos para los usuarios.

type ChallengeRepository interface {
	// InsertChallenge es una función que inserta un nuevo reto en la base de datos.
	InsertChallenge(ctx context.Context, challenge *models.Challenge) error
	// UpdateChallenge es una función que actualiza un reto en la base de datos.
	UpdateChallenge(ctx context.Context, id string, challenge *models.Challenge) error
	// DeleteChallenge es una función que elimina un reto de la base de datos por su ID.
	DeleteChallenge(ctx context.Context, id string) error
	// GetChallenges es una función que obtiene una lista de retos de la base de datos.
	GetChallenges(ctx context.Context, page int, pageSize int) ([]*models.Challenge, int, error)
	// GetChallengeById es una función que obtiene un reto de la base de datos por su ID.
	GetChallengeById(ctx context.Context, id string) (*models.Challenge, error)
	// Close es una función que cierra la conexión a la base de datos.
	CloseChallenge() error
}

// implementationChallenge es una variable que contiene la implementación de la interfaz ChallengeRepository.
var implementationChallenge ChallengeRepository

// SetChallengeRepository es una función que establece la implementación de la interfaz ChallengeRepository.
func SetChallengeRepository(repo ChallengeRepository) {
	// Establecer la implementación de la interfaz ChallengeRepository
	implementationChallenge = repo
}

// InsertChallenge es una función que inserta un nuevo reto en la base de datos.
func InsertChallenge(ctx context.Context, challenge *models.Challenge) error {
	
	// Verificar que implementationChallenge no sea nil
	if implementationChallenge == nil {
		return errors.New("implementationChallenge cannot be nil")
	}
	
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}


	// Verificar que los campos necesarios del reto no estén vacíos
	if challenge.Title == "" {
		return errors.New("challenge title cannot be empty")
	}
	if challenge.Description == "" {
		return errors.New("challenge description cannot be empty")
	}
	if challenge.Difficulty == 0 {
		return errors.New("challenge difficulty cannot be empty")
	}
	if challenge.UserID == "" {
		return errors.New("challenge user_id cannot be empty")
	}

	// Llamar a la función InsertChallenge de la implementación
	return implementationChallenge.InsertChallenge(ctx, challenge)
}


// UpdateChallenge es una función que actualiza un reto en la base de datos.
func UpdateChallenge(ctx context.Context ,id string, challenge *models.Challenge) error {
	// Verificar que implementationChallenge no sea nil
	if implementationChallenge == nil {
		return errors.New("implementationChallenge cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Verificar que el reto no sea nil
	if challenge == nil {
		return errors.New("challenge cannot be nil")
	}

	// Verificar que los campos necesarios del reto no estén vacíos
	if challenge.Title == "" {
		return errors.New("challenge title cannot be empty")
	}
	if challenge.Description == "" {
		return errors.New("challenge description cannot be empty")
	}
	if challenge.Difficulty == 0 {
		return errors.New("challenge difficulty cannot be empty")
	}
	if challenge.UserID == "" {
		return errors.New("challenge user_id cannot be empty")
	}

	// Llamar a la función UpdateChallenge de la implementación
	return implementationChallenge.UpdateChallenge(ctx, id, challenge)
}

// DeleteChallenge es una función que elimina un reto de la base de datos por su ID.
func DeleteChallenge(ctx context.Context ,id string) error {
	// Verificar que implementationChallenge no sea nil
	if implementationChallenge == nil {
		return errors.New("implementationChallenge cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Llamar a la función DeleteChallenge de la implementación
	return implementationChallenge.DeleteChallenge(ctx, id)
}


// GetChallenges es una función que obtiene una lista de retos de la base de datos.
func GetChallenges(ctx context.Context, page int, pageSize int) ([]*models.Challenge, int, error) {
	// Verificar que implementationChallenge no sea nil
	if implementationChallenge == nil {
		return nil, 0, errors.New("implementationChallenge cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return nil, 0, errors.New("context cannot be nil")
	}

	// Llamar a la función GetChallenges de la implementación
	return implementationChallenge.GetChallenges(ctx, page, pageSize)
}


// GetChallengeById es una función que obtiene un reto de la base de datos por su ID.
func GetChallengeById(ctx context.Context, id string) (*models.Challenge, error) {
	// Verificar que implementationChallenge no sea nil
	if implementationChallenge == nil {
		return nil, errors.New("implementationChallenge cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}

	// Llamar a la función GetChallengeById de la implementación
	return implementationChallenge.GetChallengeById(ctx, id)
}


// Close es una función que cierra la conexión a la base de datos.
func CloseChallenge() error {
	// Verificar que implementationChallenge no sea nil
	if implementationChallenge == nil {
		return errors.New("implementationChallenge cannot be nil")
	}

	// Llamar a la función Close de la implementación
	return implementationChallenge.CloseChallenge()
}
