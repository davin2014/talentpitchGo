package repository

import (
	"context"
	"errors"
	"talentpitchGo/models"
)

// CompanyRepository es una interfaz que define las operaciones de base de datos para las empresas.
type CompanyRepository interface {
	// InsertCompany es una función que inserta una nueva empresa en la base de datos.
	InsertCompany(ctx context.Context, company *models.Company) error
	// UpdateCompany es una función que actualiza una empresa en la base de datos.
	UpdateCompany(ctx context.Context, id string, company *models.Company) error
	// DeleteCompany es una función que elimina una empresa de la base de datos por su ID.
	DeleteCompany(ctx context.Context, id string) error
	// GetCompanies es una función que obtiene una lista de empresas de la base de datos.
	GetCompanies(ctx context.Context, page int, pageSize int) ([]*models.Company, int, error)
	// GetCompanyById es una función que obtiene una empresa de la base de datos por su ID.
	GetCompanyById(ctx context.Context, id string) (*models.Company, error)
	// CloseCompany es una función que cierra la conexión a la base de datos.
	CloseCompany() error
}

// implementationCompany es una variable que contiene la implementación de la interfaz CompanyRepository.
var implementationCompany CompanyRepository

// SetCompanyRepository es una función que establece la implementación de la interfaz CompanyRepository.
func SetCompanyRepository(repo CompanyRepository) {
	// Establecer la implementación de la interfaz CompanyRepository
	implementationCompany = repo
}

// InsertCompany es una función que inserta una nueva empresa en la base de datos.
func InsertCompany(ctx context.Context, company *models.Company) error {
	// Verificar que implementationCompany no sea nil
	if implementationCompany == nil {
		return errors.New("implementationCompany cannot be nil")
	}
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	// Verificar que los campos necesarios de la empresa no estén vacíos
	if company.Name == "" {
		return errors.New("company name cannot be empty")
	}
	if company.Location == "" {
		return errors.New("company location cannot be empty")
	}
	if company.Industry == "" {
		return errors.New("company industry cannot be empty")
	}
	if company.UserID == "" {
		return errors.New("company user_id cannot be empty")
	}
	// Insertar la empresa en la base de datos
	return implementationCompany.InsertCompany(ctx, company)
}

// UpdateCompany es una función que actualiza una empresa en la base de datos.
func UpdateCompany(ctx context.Context, id string, company *models.Company) error {
	// Verificar que implementationCompany no sea nil
	if implementationCompany == nil {
		return errors.New("implementationCompany cannot be nil")
	}
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	// Verificar que los campos necesarios de la empresa no estén vacíos
	if company.Name == "" {
		return errors.New("company name cannot be empty")
	}
	if company.Location == "" {
		return errors.New("company location cannot be empty")
	}
	if company.Industry == "" {
		return errors.New("company industry cannot be empty")
	}
	if company.UserID == "" {
		return errors.New("company user_id cannot be empty")
	}
	// Actualizar la empresa en la base de datos
	return implementationCompany.UpdateCompany(ctx, id, company)
}

// DeleteCompany es una función que elimina una empresa de la base de datos por su ID.
func DeleteCompany(ctx context.Context, id string) error {
	// Verificar que implementationCompany no sea nil
	if implementationCompany == nil {
		return errors.New("implementationCompany cannot be nil")
	}
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	// Eliminar la empresa de la base de datos
	return implementationCompany.DeleteCompany(ctx, id)
}


// GetCompanies es una función que obtiene una lista de empresas de la base de datos.
func GetCompanies(ctx context.Context, page int, pageSize int) ([]*models.Company, int, error) {
	// Verificar que implementationCompany no sea nil
	if implementationCompany == nil {
		return nil, 0, errors.New("implementationCompany cannot be nil")
	}
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return nil, 0, errors.New("context cannot be nil")
	}
	// Obtener una lista de empresas de la base de datos
	return implementationCompany.GetCompanies(ctx, page, pageSize)
}

// GetCompanyById es una función que obtiene una empresa de la base de datos por su ID.
func GetCompanyById(ctx context.Context, id string) (*models.Company, error) {
	// Verificar que implementationCompany no sea nil
	if implementationCompany == nil {
		return nil, errors.New("implementationCompany cannot be nil")
	}
	// Verificar que el contexto no sea nil
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	// Obtener la empresa de la base de datos
	return implementationCompany.GetCompanyById(ctx, id)
}

// CloseCompany es una función que cierra la conexión a la base de datos.
func CloseCompany() error {
	// Verificar que implementationCompany no sea nil
	if implementationCompany == nil {
		return errors.New("implementationCompany cannot be nil")
	}
	// Cerrar la conexión a la base de datos
	return implementationCompany.CloseCompany()
}

	