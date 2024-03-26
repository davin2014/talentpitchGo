
CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(255) PRIMARY KEY,
  FULLName VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);  

CREATE TABLE IF NOT EXISTS challenges (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    difficulty INT,
    user_id VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS companies (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    image_path VARCHAR(255) NULL,
    location VARCHAR(255),
    industry VARCHAR(255),
    user_id VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);
 
CREATE TABLE IF NOT EXISTS programs (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    start_date DATE,
    end_date DATE,
    user_id VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);




 
CREATE TABLE IF NOT EXISTS program_participants (
    id VARCHAR(255) PRIMARY KEY,
    program_id VARCHAR(255),
    entity_type VARCHAR(50), -- Indicar el tipo de entidad: 'user', 'challenge', 'company'
    entity_id VARCHAR(255), -- ID de la entidad participante
    FOREIGN KEY (program_id) REFERENCES programs(id)
);
 


