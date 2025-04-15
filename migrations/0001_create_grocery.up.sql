CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE status_enum AS ENUM ('active', 'inactive', 'pending', 'review', 'cancelled', 'deleted');


CREATE TABLE grocery (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    address TEXT,
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20),
    email VARCHAR(100),
    position TEXT,
    opening_hours VARCHAR(100),
    closing_hours VARCHAR(100),
    is_open BOOLEAN DEFAULT FALSE,
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token VARCHAR(255),
    verification_token_expiry TIMESTAMP,
    status status_enum NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);


CREATE INDEX idx_grocery_name ON grocery (name);
CREATE INDEX idx_grocery_city ON grocery (city);
CREATE INDEX idx_phone_number ON grocery (phone_number);
CREATE INDEX idx_email ON grocery (email);