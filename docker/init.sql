-- Create user superman if not exists
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'superman') THEN

      CREATE USER superman WITH PASSWORD 'abc@123@xyz';
   END IF;
END
$do$;

-- Grant all privileges on all tables in grocery_db to superman
GRANT ALL PRIVILEGES ON DATABASE grocery_store TO superman;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO superman;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO superman;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO superman;

-- Grant future tables permissions
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO superman;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO superman;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON FUNCTIONS TO superman;

GRANT ALL ON SCHEMA public TO superman;
GRANT ALL ON ALL TABLES IN SCHEMA public TO superman;
ALTER USER superman WITH CREATEDB;