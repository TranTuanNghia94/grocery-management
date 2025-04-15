CREATE TYPE  price_change_reason AS ENUM ('initial', 'promotion', 'seasonal', 'supplier_change', 'market_adjustment', 'cost_increase', 'cost_decrease', 'competitive_adjustment', 'other');

CREATE TABLE price_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
    old_price DECIMAL(15, 0) NULL,
    new_price DECIMAL(15, 0) NOT NULL,
    effective_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP NULL,
    is_current BOOLEAN NOT NULL DEFAULT TRUE,
    vendor_id UUID NULL,
    contact_id UUID NULL,
    change_reason price_change_reason NOT NULL DEFAULT 'initial',
    notes TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    FOREIGN KEY (vendor_id) REFERENCES vendors(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (contact_id) REFERENCES contacts(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Create indexes for the price_history table
CREATE INDEX idx_price_history_product_id ON price_history (product_id);
CREATE INDEX idx_price_history_effective_date ON price_history (effective_date);
CREATE INDEX idx_price_history_is_current ON price_history (is_current);