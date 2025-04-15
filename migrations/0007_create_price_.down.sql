-- Drop indexes
DROP INDEX IF EXISTS idx_price_history_product_id;
DROP INDEX IF EXISTS idx_price_history_effective_date;
DROP INDEX IF EXISTS idx_price_history_is_current;

-- Drop the price_history table
DROP TABLE IF EXISTS price_history;

-- Drop the price_change_reason enum type
DROP TYPE IF EXISTS price_change_reason;
