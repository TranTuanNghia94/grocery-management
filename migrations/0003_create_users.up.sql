CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grocery_id UUID NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone_number VARCHAR(20),
    status status_enum NOT NULL DEFAULT 'active',
    role_id UUID REFERENCES roles(id),
    is_active BOOLEAN DEFAULT TRUE,
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,

    CONSTRAINT fk_grocery FOREIGN KEY (grocery_id) REFERENCES grocery(id),
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id),
    CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES users(id),
    CONSTRAINT fk_deleted_by FOREIGN KEY (deleted_by) REFERENCES users(id)
);

CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_username ON users (username);

CREATE INDEX idx_users_phone_number ON users (phone_number);
CREATE INDEX idx_users_grocery_id ON users (grocery_id);
