// El paquete database contiene las funciones para interactuar con la base de datos.
package database

import (
	"errors"
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"talentpitchGo/models"
)

// PostgresRepositoy es una estructura que contiene la conexión a la base de datos.
type PostgresRepositoy struct {
	db *sql.DB
}


// NewPostgresRepository es una función que crea una nueva conexión a la base de datos
// y devuelve un nuevo repositorio Postgres.
func NewPostgresRepository(url string) (*PostgresRepositoy, error ){
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepositoy{db: db}, nil
}


// InsertUser es una función que inserta un nuevo usuario en la base de datos.
func (p *PostgresRepositoy) InsertUser(ctx context.Context ,user *models.User) error {
	// Verificar que la base de datos está disponible
    _, err := p.db.ExecContext(ctx, "SELECT 1")
    if err != nil {
        return errors.New("database is not available")
    }

	// Verificar que el repositorio no sea nil
    if p == nil {
        return errors.New("repository cannot be nil")
    }

    // Verificar que la base de datos no sea nil
    if p.db == nil {
        return errors.New("database cannot be nil")
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
    if user.Id == "" {
        return errors.New("user id cannot be empty")
    }
    if user.Fullname == "" {
        return errors.New("user fullname cannot be empty")
    }
    if user.Email == "" {
        return errors.New("user email cannot be empty")
    }
    if user.Password == "" {
        return errors.New("user password cannot be empty")
    }
	// Insertar un nuevo usuario en la base de datos
	_, err = p.db.ExecContext(ctx, "INSERT INTO users (id, fullname, email, password) VALUES ($1, $2, $3, $4)",user.Id, user.Fullname, user.Email, user.Password)
	return err
}


// Update actualiza un elemento en la base de datos
func (p *PostgresRepositoy) UpdateUser(ctx context.Context ,id string, user *models.User)  error {
    // Ejecutar la consulta de actualización
	_, err := p.db.ExecContext(ctx, "UPDATE users SET fullname = $1, email = $2 WHERE id = $3", user.Fullname, user.Email, id)
	// Manejar el error si existe
	if err != nil {
		log.Fatal("Error closing database connection: ", err)
	}
	// Cerrar la conexión a la base de datos
	defer func() {
		if err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
	}()
	// Devolver nil si la actualización fue exitosa
    return  nil
}


// Delete elimina un elemento de la base de datos
func (p *PostgresRepositoy) DeleteUser(ctx context.Context, id string) error {
    
	log.Println("Deleting user with id: ", id)
    // Ejecutar la consulta de eliminación
    _, err := p.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1 ", id)
    // Manejar el error si existe
    if err != nil {
        return err
    }
    
    // Devolver nil si la eliminación fue exitosa
    return nil
}


// GetUsers obtiene usuarios de la base de datos con paginación
func (p *PostgresRepositoy) GetUsers(ctx context.Context, page int, pageSize int) ([]*models.User, int, error) {
	var users []*models.User
    // Comprobar si la página y el tamaño de la página son válidos
    if page < 1 || pageSize < 1 {
        return nil, 0, errors.New("invalid page number or page size")
    }

    // Calcular el offset
    offset := (page - 1) * pageSize

    // Ejecutar la consulta para obtener usuarios
    rows, err := p.db.QueryContext(ctx, "SELECT id, fullname, email, password FROM users ORDER BY id LIMIT $1 OFFSET $2", pageSize, offset)
    if err != nil {
        return nil, 0, err
    }
    defer rows.Close()

    // Iterar sobre los resultados
    
    for rows.Next() {
		var user = models.User{}
        if err = rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password); err != nil {
            return nil, 0, err
        }
        users = append(users, &user)
    }

    // Comprobar si hubo errores durante la iteración
    if err = rows.Err(); err != nil {
        return nil, 0, err
    }

    // Ejecutar la consulta para contar el número total de usuarios
    row := p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users")
    var count int
    err = row.Scan(&count)
    if err != nil {
        return nil, 0, err
    }
    // Devolver los usuarios y el conteo
    return users, count, nil
}


// GetUserById es una función que obtiene un usuario de la base de datos por su ID.
func (p *PostgresRepositoy) GetUserById(ctx context.Context, id string) (*models.User, error) {
	// Crear una nueva estructura de usuario
	var user = models.User{}
	// Obtener un usuario de la base de datos por su ID
	rows, err := p.db.QueryContext(ctx, "SELECT id, fullname, email FROM users WHERE id = $1", id)
	// Manejar el error si existe
	if err != nil {
		return nil, err
	}
	// Cerrar la conexión a la base de datos
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
	}()
	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Fullname, &user.Email); err == nil {
			return &user, nil
		}
	}

	// Iterar sobre los resultados de la consulta
	if !rows.Next() {
		return nil, fmt.Errorf("no user found with id %s", id)
	}
    // Manejar el error si existe
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Devolver el usuario
	return &user, nil
}


// GetUserByEmail es una función que obtiene un usuario de la base de datos por su email.
func (p *PostgresRepositoy) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	// Crear una nueva estructura de usuario
	var user = models.User{}
	// Obtener un usuario de la base de datos por su email
	rows, err := p.db.QueryContext(ctx, "SELECT id, fullname, email, password FROM users WHERE email = $1", email)
	// Manejar el error si existe	
    if err != nil {
		return nil, err
	}
	// Cerrar la conexión a la base de datos
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
	}()
	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Fullname, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
    // Manejar el error si existe
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Devolver el usuario
	return &user, nil
}


// Close es una función que cierra la conexión a la base de datos.
func (p *PostgresRepositoy) Close() error {
	return p.db.Close()
}


//********************************************************************************************************************
//************************************************************* CHALLENGE ************************************************
//********************************************************************************************************************

// InsertChallenge es una función que inserta un nuevo reto en la base de datos.
func (p *PostgresRepositoy) InsertChallenge(ctx context.Context, challenge *models.Challenge) error {
	// Verificar que la base de datos está disponible
	_, err := p.db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return errors.New("database is not available")
	}

	// Verificar que el repositorio no sea nil
	if p == nil {
		return errors.New("repository cannot be nil")
	}

	// Verificar que la base de datos no sea nil
	if p.db == nil {
		return errors.New("database cannot be nil")
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

	// Insertar un nuevo reto en la base de datos
	_, err = p.db.ExecContext(ctx, "INSERT INTO challenges (id, title, description, difficulty, user_id) VALUES ($1, $2, $3, $4, $5)",challenge.Id, challenge.Title, challenge.Description, challenge.Difficulty, challenge.UserID)
	return err
}


// UpdateChallenge es una función que actualiza un reto en la base de datos.
func (p *PostgresRepositoy) UpdateChallenge(ctx context.Context ,id string, challenge *models.Challenge) error {
	// Verificar que la base de datos está disponible
	_, err := p.db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return errors.New("database is not available")
	}

	// Verificar que el repositorio no sea nil
	if p == nil {
		return errors.New("repository cannot be nil")
	}

	// Verificar que la base de datos no sea nil
	if p.db == nil {
		return errors.New("database cannot be nil")
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

	// Actualizar un reto en la base de datos
	_, err = p.db.ExecContext(ctx, "UPDATE challenges SET title = $1, description = $2, difficulty = $3, user_id = $4 WHERE id = $5", challenge.Title, challenge.Description, challenge.Difficulty, challenge.UserID, id)
	return err
}


// DeleteChallenge es una función que elimina un reto de la base de datos.
func (p *PostgresRepositoy) DeleteChallenge(ctx context.Context, id string) error {
	// Verificar que la base de datos está disponible
	_, err := p.db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return errors.New("database is not available")
	}

	// Verificar que el repositorio no sea nil
	if p == nil {
		return errors.New("repository cannot be nil")
	}

	// Verificar que la base de datos no sea nil
	if p.db == nil {
		return errors.New("database cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Eliminar un reto de la base de datos
	_, err = p.db.ExecContext(ctx, "DELETE FROM challenges WHERE id = $1", id)
	return err
}


// GetChallenges es una función que obtiene retos de la base de datos con paginación.
func (p *PostgresRepositoy) GetChallenges(ctx context.Context, page int, pageSize int) ([]*models.Challenge, int, error) {
	var challenges []*models.Challenge
	// Comprobar si la página y el tamaño de la página son válidos
	if page < 1 || pageSize < 1 {
		return nil, 0, errors.New("invalid page number or page size")
	}

	// Calcular el offset
	offset := (page - 1) * pageSize

	// Ejecutar la consulta para obtener retos
	rows, err := p.db.QueryContext(ctx, "SELECT id, title, description, difficulty, user_id FROM challenges ORDER BY id LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterar sobre los resultados
	for rows.Next() {
		var challenge = models.Challenge{}
		if err = rows.Scan(&challenge.Id, &challenge.Title, &challenge.Description, &challenge.Difficulty, &challenge.UserID); err != nil {
			return nil, 0, err
		}
		challenges = append(challenges, &challenge)
	}

	// Comprobar si hubo errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Ejecutar la consulta para contar el número total de retos
	row := p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM challenges")
	var count int
	err = row.Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	// Devolver los retos y el conteo
	return challenges, count, nil
}


// GetChallengeById es una función que obtiene un reto de la base de datos por su ID.
func (p *PostgresRepositoy) GetChallengeById(ctx context.Context, id string) (*models.Challenge, error) {
	// Crear una nueva estructura de reto
	var challenge = models.Challenge{}
	// Obtener un reto de la base de datos por su ID
	rows, err := p.db.QueryContext(ctx, "SELECT id, title, description, difficulty, user_id FROM challenges WHERE id = $1", id)
	
	// Manejar el error si existe
	if err != nil {
		return nil, err
	}
	// Cerrar la conexión a la base de datos
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
	}()
	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		if err = rows.Scan(&challenge.Id, &challenge.Title, &challenge.Description, &challenge.Difficulty, &challenge.UserID); err == nil {
			return &challenge, nil
		}
	}

	// Iterar sobre los resultados de la consulta
	if !rows.Next() {
		return nil, fmt.Errorf("no user found with id %s", id)
	}
    // Manejar el error si existe
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	// Devolver el reto
	return &challenge, nil
}


// Close es una función que cierra la conexión a la base de datos.
func (p *PostgresRepositoy) CloseChallenge() error {
	return p.db.Close()
}


//********************************************************************************************************************
//************************************************************* COMPANY ************************************************
//********************************************************************************************************************

// InsertCompany es una función que inserta una nueva empresa en la base de datos.
func (p *PostgresRepositoy) InsertCompany(ctx context.Context, company *models.Company) error {
	// Verificar que la base de datos está disponible
	_, err := p.db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return errors.New("database is not available")
	}

	// Verificar que el repositorio no sea nil
	if p == nil {
		return errors.New("repository cannot be nil")
	}

	// Verificar que la base de datos no sea nil
	if p.db == nil {
		return errors.New("database cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Verificar que la empresa no sea nil
	if company == nil {
		return errors.New("company cannot be nil")
	}

	// Verificar que los campos necesarios de la empresa no estén vacíos
	if company.Name == "" {
		return errors.New("company name cannot be empty")
	}
	if company.ImagePath == "" {
		return errors.New("company image_path cannot be empty")
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

	// Insertar una nueva empresa en la base de datos
	_, err = p.db.ExecContext(ctx, "INSERT INTO companies (id, name, image_path, location, industry, user_id) VALUES ($1, $2, $3, $4, $5, $6)",company.Id, company.Name, company.ImagePath, company.Location, company.Industry, company.UserID)
	return err
}


// UpdateCompany es una función que actualiza una empresa en la base de datos.
func (p *PostgresRepositoy) UpdateCompany(ctx context.Context, id string, company *models.Company) error {
	// Verificar que la base de datos está disponible
	_, err := p.db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return errors.New("database is not available")
	}

	// Verificar que el repositorio no sea nil
	if p == nil {
		return errors.New("repository cannot be nil")
	}

	// Verificar que la base de datos no sea nil
	if p.db == nil {
		return errors.New("database cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Verificar que la empresa no sea nil
	if company == nil {
		return errors.New("company cannot be nil")
	}

	// Verificar que los campos necesarios de la empresa no estén vacíos
	if company.Name == "" {
		return errors.New("company name cannot be empty")
	}
	if company.ImagePath == "" {
		return errors.New("company image_path cannot be empty")
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

	// Actualizar una empresa en la base de datos
	_, err = p.db.ExecContext(ctx, "UPDATE companies SET name = $1, image_path = $2, location = $3, industry = $4, user_id = $5 WHERE id = $6", company.Name, company.ImagePath, company.Location, company.Industry, company.UserID, id)
	return err
}

// DeleteCompany es una función que elimina una empresa de la base de datos por su ID.
func (p *PostgresRepositoy) DeleteCompany(ctx context.Context, id string) error {
	// Verificar que la base de datos está disponible
	_, err := p.db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return errors.New("database is not available")
	}

	// Verificar que el repositorio no sea nil
	if p == nil {
		return errors.New("repository cannot be nil")
	}

	// Verificar que la base de datos no sea nil
	if p.db == nil {
		return errors.New("database cannot be nil")
	}

	// Verificar que el contexto no sea nil
	if ctx == nil {
		return errors.New("context cannot be nil")
	}

	// Eliminar una empresa de la base de datos
	_, err = p.db.ExecContext(ctx, "DELETE FROM companies WHERE id = $1", id)
	return err
}

// GetCompanies es una función que obtiene empresas de la base de datos con paginación.
func (p *PostgresRepositoy) GetCompanies(ctx context.Context, page int, pageSize int) ([]*models.Company, int, error) {
	var companies []*models.Company
	// Comprobar si la página y el tamaño de la página son válidos
	if page < 1 || pageSize < 1 {
		return nil, 0, errors.New("invalid page number or page size")
	}

	// Calcular el offset
	offset := (page - 1) * pageSize

	// Ejecutar la consulta para obtener empresas
	rows, err := p.db.QueryContext(ctx, "SELECT id, name, image_path, location, industry, user_id FROM companies ORDER BY id LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterar sobre los resultados
	for rows.Next() {
		var company = models.Company{}
		if err = rows.Scan(&company.Id, &company.Name, &company.ImagePath, &company.Location, &company.Industry, &company.UserID); err != nil {
			return nil, 0, err
		}
		companies = append(companies, &company)
	}

	// Comprobar si hubo errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Ejecutar la consulta para contar el número total de empresas
	row := p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM companies")
	var count int
	err = row.Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	// Devolver las empresas y el conteo
	return companies, count, nil
}

// GetCompanyById es una función que obtiene una empresa de la base de datos por su ID.
func (p *PostgresRepositoy) GetCompanyById(ctx context.Context, id string) (*models.Company, error) {
	// Crear una nueva estructura de empresa
	var company = models.Company{}
	// Obtener una empresa de la base de datos por su ID
	rows, err := p.db.QueryContext(ctx, "SELECT id, name, image_path, location, industry, user_id FROM companies WHERE id = $1", id)
	// Manejar el error si existe
	if err != nil {
		return nil, err
	}
	// Cerrar la conexión a la base de datos
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
	}()
	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		if err = rows.Scan(&company.Id, &company.Name, &company.ImagePath, &company.Location, &company.Industry, &company.UserID); err == nil {
			return &company, nil
		}
	}

	// Iterar sobre los resultados de la consulta
	if !rows.Next() {
		return nil, fmt.Errorf("no user found with id %s", id)
	}
	// Manejar el error si existe
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Devolver la empresa
	return &company, nil
}

// CloseCompany es una función que cierra la conexión a la base de datos.
func (p *PostgresRepositoy) CloseCompany() error {
	return p.db.Close()
}

