-- Drop indexes on contacts table
DROP INDEX IF EXISTS idx_contacts_name;
DROP INDEX IF EXISTS idx_contacts_email;
DROP INDEX IF EXISTS idx_contacts_phone_number;
DROP INDEX IF EXISTS idx_contacts_status;

-- Drop contacts table
DROP TABLE IF EXISTS contacts;

-- Drop indexes on vendors table
DROP INDEX IF EXISTS idx_vendors_name;
DROP INDEX IF EXISTS idx_vendors_code;
DROP INDEX IF EXISTS idx_vendors_email;
DROP INDEX IF EXISTS idx_vendors_phone_number;
DROP INDEX IF EXISTS idx_vendors_status;

-- Drop vendors table
DROP TABLE IF EXISTS vendors;
