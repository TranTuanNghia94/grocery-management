CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE  inventory_status AS ENUM ('available', 'out_of_stock', 'not_available', 'need_restock', 'backorder', 'damaged', 'expired');

CREATE TYPE  inventory_transaction_type AS ENUM ('add', 'remove', 'adjustment', 'transfer', 'return', 'write_off', 'damaged', 'expired', 'other', 'update_price');


CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grocery_id UUID NOT NULL REFERENCES grocery(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    sku VARCHAR(80) NOT NULL,
    barcode VARCHAR(100) NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    product_name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    unit_price DECIMAL(10, 2) NOT NULL,
    expiry_date DATE,
    reorder_level INTEGER DEFAULT 10,
    in_stock BOOLEAN DEFAULT TRUE,
    inventory_status inventory_status NOT NULL DEFAULT 'available',
    created_by UUID,
    updated_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inventory_grocery_id ON inventory (grocery_id);
CREATE INDEX idx_inventory_product_id ON inventory (product_id);
CREATE INDEX idx_inventory_sku ON inventory (sku);
CREATE INDEX idx_inventory_barcode ON inventory (barcode);

CREATE INDEX idx_inventory_category_id ON inventory (category_id);

CREATE TABLE inventory_transaction (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    grocery_id UUID NOT NULL REFERENCES grocery(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    sku VARCHAR(80) NOT NULL,
    barcode VARCHAR(100) NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    old_quantity INTEGER NOT NULL,
    new_quantity INTEGER NOT NULL,
    quantity_change INTEGER NOT NULL,
    unit_price DECIMAL(15, 0) NOT NULL,
    transaction_type inventory_transaction_type NOT NULL,
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    updated_by UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inventory_transaction_grocery_id ON inventory_transaction (grocery_id);
CREATE INDEX idx_inventory_transaction_product_id ON inventory_transaction (product_id);
CREATE INDEX idx_inventory_transaction_sku ON inventory_transaction (sku);