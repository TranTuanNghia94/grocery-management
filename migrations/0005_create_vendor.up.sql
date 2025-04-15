CREATE TABLE vendors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(200) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    phone_number VARCHAR(20),
    email VARCHAR(100),
    website VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100),
    address TEXT,
    status status_enum NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID
);

-- Create indexes for the vendor table
CREATE INDEX idx_vendors_name ON vendors (name);
CREATE INDEX idx_vendors_code ON vendors (code);
CREATE INDEX idx_vendors_email ON vendors (email);
CREATE INDEX idx_vendors_phone_number ON vendors (phone_number);
CREATE INDEX idx_vendors_status ON vendors (status);


CREATE TABLE contacts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    vendor_id UUID REFERENCES vendors(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20),
    email VARCHAR(100),
    address TEXT,
    designation VARCHAR(100),
    department VARCHAR(100),
    date_of_birth DATE,
    first_name VARCHAR(150),
    last_name VARCHAR(150),
    position VARCHAR(100),
    status status_enum NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID
);

-- Create indexes for the contact_person table
CREATE INDEX idx_contacts_name ON contacts (name);
CREATE INDEX idx_contacts_email ON contacts (email);
CREATE INDEX idx_contacts_phone_number ON contacts (phone_number);
CREATE INDEX idx_contacts_status ON contacts (status);