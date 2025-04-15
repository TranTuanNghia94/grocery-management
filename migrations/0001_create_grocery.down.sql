-- Drop indexes
DROP INDEX IF EXISTS idx_grocery_name;
DROP INDEX IF EXISTS idx_grocery_city;
DROP INDEX IF EXISTS idx_phone_number;
DROP INDEX IF EXISTS idx_email;

-- Drop table
DROP TABLE IF EXISTS grocery;

-- Drop enum type
DROP TYPE IF EXISTS status_enum;
