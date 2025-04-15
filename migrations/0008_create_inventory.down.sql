-- Drop indexes first
DROP INDEX IF EXISTS idx_inventory_transaction_sku;
DROP INDEX IF EXISTS idx_inventory_transaction_product_id;
DROP INDEX IF EXISTS idx_inventory_transaction_grocery_id;

DROP INDEX IF EXISTS idx_inventory_category_id;
DROP INDEX IF EXISTS idx_inventory_barcode;
DROP INDEX IF EXISTS idx_inventory_sku;
DROP INDEX IF EXISTS idx_inventory_product_id;
DROP INDEX IF EXISTS idx_inventory_grocery_id;

-- Drop tables in reverse order of creation (respecting foreign key constraints)
DROP TABLE IF EXISTS inventory_transaction;
DROP TABLE IF EXISTS inventory;

-- Drop custom types
DROP TYPE IF EXISTS inventory_transaction_type;
DROP TYPE IF EXISTS inventory_status;
