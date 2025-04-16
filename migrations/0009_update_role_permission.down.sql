-- Rollback the changes made in 0009_update_role_permission.up.sql
ALTER TABLE role_permissions DROP COLUMN description;