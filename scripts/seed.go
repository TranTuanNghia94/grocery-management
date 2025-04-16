package main

import (
	"database/sql"
	"grocery-management/internal/models"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=grocery_store sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	seedRoles(db)
	seedPermissions(db)
	seedRolePermissions(db)
	seedGrocery(db)
}

func getStringVlue(s string) *string {
	return &s
}

var mapRoleToID = map[string]string{}

func seedRoles(db *sql.DB) {
	log.Println("Starting to seed roles...")

	roles := []models.Role{
		{
			Name:        "admin",
			Description: getStringVlue("Full system access"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "manager",
			Description: getStringVlue("Department management access"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "staff",
			Description: getStringVlue("Basic operational access"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "guest",
			Description: getStringVlue("Read-only access"),
			Status:      getStringVlue("active"),
		},
	}

	for _, role := range roles {
		result, err := db.Query("INSERT INTO roles (name, description, status) VALUES ($1, $2, $3) RETURNING id",
			role.Name, role.Description, role.Status)
		if err != nil {
			log.Printf("Error seeding role %s: %v", role.Name, err)
		} else {
			var id string
			if result.Next() {
				result.Scan(&id)
				log.Printf("Role added: %s with ID: %s", role.Name, id)
				mapRoleToID[role.Name] = id
			}
		}
	}

	log.Println("Mapping roles to IDs:", mapRoleToID)

	log.Println("Seeding roles completed.")
}

var mapPermissionToID = map[string]string{}

func seedPermissions(db *sql.DB) {
	log.Println("Start seeding permissions...")

	permissions := []models.Permission{
		{
			Name:        "write_all",
			Description: getStringVlue("Create all items"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "view_all",
			Description: getStringVlue("Read all items"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "update_all",
			Description: getStringVlue("Update all items"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "delete_all",
			Description: getStringVlue("Delete all items"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "manage_users",
			Description: getStringVlue("Add/remove users"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "view_reports",
			Description: getStringVlue("Access reporting features"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "manage_inventory",
			Description: getStringVlue("Control inventory levels"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "view_products",
			Description: getStringVlue("View products"),
			Status:      getStringVlue("active"),
		},
		{
			Name:        "manage_products",
			Description: getStringVlue("Manage products"),
			Status:      getStringVlue("active"),
		},
	}

	for _, perm := range permissions {
		result, err := db.Query("INSERT INTO permissions (name, description, status) VALUES ($1, $2, $3) RETURNING id",
			perm.Name, perm.Description, perm.Status)
		if err != nil {
			log.Printf("Error seeding permission %s: %v", perm.Name, err)
		} else {
			var id string
			if result.Next() {
				result.Scan(&id)
				log.Printf("Permission added: %s with ID: %s", perm.Name, id)
				mapPermissionToID[perm.Name] = id
			}
		}
	}

	log.Println("Mapping permissions to IDs:", mapPermissionToID)

	log.Println("Seeding permissions completed.")
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func seedRolePermissions(db *sql.DB) {
	log.Println("Starting to seed role permissions...")

	rolePermissions := []models.RolePermission{}

	for k, v := range mapPermissionToID {

		if contains(k, "all") {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleID:       mapRoleToID["admin"],
				PermissionID: v,
				Description:  getStringVlue("admin_" + k),
			})
		} else if contains(k, "manage") {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleID:       mapRoleToID["manager"],
				PermissionID: v,
				Description:  getStringVlue("manager_" + k),
			})
		} else {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleID:       mapRoleToID["staff"],
				PermissionID: v,
				Description:  getStringVlue("staff_" + k),
			})
		}
	}

	for _, rolePermission := range rolePermissions {
		_, err := db.Exec("INSERT INTO role_permissions (role_id, permission_id, Description) VALUES ($1, $2, $3)",
			rolePermission.RoleID, rolePermission.PermissionID, rolePermission.Description)
		if err != nil {
			log.Printf("Error seeding role permission %s: %v", rolePermission.PermissionID, err)
		} else {
			log.Printf("Role permission added: Role ID: %s, Permission ID: %s", rolePermission.RoleID, rolePermission.PermissionID)
		}
	}
	log.Println("Seeding role permissions completed.")
	log.Println("Seeding completed.")

}

func seedGrocery(db *sql.DB) {
	log.Println("Starting to seed grocery...")

	grocery := models.Grocery{
		Name:         "Grocery Store Local",
		Description:  getStringVlue("A local grocery store"),
		Status:       getStringVlue("active"),
		City:         getStringVlue("Ho Chi Minh"),
		Country:      getStringVlue("Vietnam"),
		Address:      getStringVlue("123 Main St"),
		PostalCode:   getStringVlue("700000"),
		PhoneNumber:  getStringVlue("123456789"),
		OpeningHours: getStringVlue("9:00 AM - 9:00 PM"),
	}

	_, err := db.Exec("INSERT INTO grocery (name, description, status, city, country, address, postal_code, phone_number, opening_hours) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		grocery.Name, grocery.Description, grocery.Status, grocery.City, grocery.Country, grocery.Address, grocery.PostalCode, grocery.PhoneNumber, grocery.OpeningHours)
	if err != nil {
		log.Printf("Error seeding grocery %s: %v", grocery.Name, err)
	} else {
		log.Printf("Grocery added: %s", grocery.Name)
	}

	log.Println("Seeding grocery completed.")
}
