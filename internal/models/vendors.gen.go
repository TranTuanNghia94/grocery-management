// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVendor = "vendors"

// Vendor mapped from table <vendors>
type Vendor struct {
	ID          *string        `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string         `gorm:"column:name;type:character varying(200);not null" json:"name"`
	Code        string         `gorm:"column:code;type:character varying(50);not null" json:"code"`
	Description *string        `gorm:"column:description;type:text" json:"description"`
	PhoneNumber *string        `gorm:"column:phone_number;type:character varying(20)" json:"phone_number"`
	Email       *string        `gorm:"column:email;type:character varying(100)" json:"email"`
	Website     *string        `gorm:"column:website;type:character varying(255)" json:"website"`
	City        *string        `gorm:"column:city;type:character varying(100)" json:"city"`
	State       *string        `gorm:"column:state;type:character varying(100)" json:"state"`
	Country     *string        `gorm:"column:country;type:character varying(100)" json:"country"`
	Address     *string        `gorm:"column:address;type:text" json:"address"`
	Status      *string        `gorm:"column:status;type:status_enum;not null;default:active" json:"status"`
	CreatedAt   *time.Time     `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp without time zone" json:"deleted_at"`
	CreatedBy   *string        `gorm:"column:created_by;type:uuid" json:"created_by"`
	UpdatedBy   *string        `gorm:"column:updated_by;type:uuid" json:"updated_by"`
	DeletedBy   *string        `gorm:"column:deleted_by;type:uuid" json:"deleted_by"`
}

// TableName Vendor's table name
func (*Vendor) TableName() string {
	return TableNameVendor
}
