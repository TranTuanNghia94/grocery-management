// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           *string        `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	GroceryID    string         `gorm:"column:grocery_id;type:uuid;not null" json:"grocery_id"`
	Username     string         `gorm:"column:username;type:character varying(50);not null" json:"username"`
	Email        string         `gorm:"column:email;type:character varying(100);not null" json:"email"`
	PasswordHash string         `gorm:"column:password_hash;type:character varying(255);not null" json:"password_hash"`
	FirstName    *string        `gorm:"column:first_name;type:character varying(50)" json:"first_name"`
	LastName     *string        `gorm:"column:last_name;type:character varying(50)" json:"last_name"`
	PhoneNumber  *string        `gorm:"column:phone_number;type:character varying(20)" json:"phone_number"`
	Status       *string        `gorm:"column:status;type:status_enum;not null;default:active" json:"status"`
	RoleID       *string        `gorm:"column:role_id;type:uuid" json:"role_id"`
	IsActive     *bool          `gorm:"column:is_active;type:boolean;default:true" json:"is_active"`
	LastLogin    *time.Time     `gorm:"column:last_login;type:timestamp without time zone" json:"last_login"`
	CreatedAt    *time.Time     `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    *time.Time     `gorm:"column:updated_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp without time zone" json:"deleted_at"`
	CreatedBy    *string        `gorm:"column:created_by;type:uuid" json:"created_by"`
	UpdatedBy    *string        `gorm:"column:updated_by;type:uuid" json:"updated_by"`
	DeletedBy    *string        `gorm:"column:deleted_by;type:uuid" json:"deleted_by"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
