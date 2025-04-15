CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    sku VARCHAR(80) UNIQUE,
    barcode VARCHAR(100),
    category_id UUID REFERENCES categories(id),
    price DECIMAL(15, 0) NOT NULL,
    cost_price DECIMAL(15, 0),
    tax_rate DECIMAL(5, 2) DEFAULT 0,
    weight DECIMAL(8, 2),
    weight_unit VARCHAR(10),
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    status status_enum NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_name ON products (name);
CREATE INDEX idx_products_sku ON products (sku);
CREATE INDEX idx_products_category_id ON products (category_id);
