-- Drop indexes first
DROP INDEX IF EXISTS idx_roles_name;
DROP INDEX IF EXISTS idx_role_permissions_status;
DROP INDEX IF EXISTS idx_role_permissions_permission_id;
DROP INDEX IF EXISTS idx_role_permissions_role_id;

-- Drop tables in reverse order to avoid foreign key constraint issues
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
