// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePriceHistory = "price_history"

// PriceHistory mapped from table <price_history>
type PriceHistory struct {
	ID            *string        `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProductID     *string        `gorm:"column:product_id;type:uuid" json:"product_id"`
	OldPrice      *float64       `gorm:"column:old_price;type:numeric(15,0)" json:"old_price"`
	NewPrice      float64        `gorm:"column:new_price;type:numeric(15,0);not null" json:"new_price"`
	EffectiveDate *time.Time     `gorm:"column:effective_date;type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP" json:"effective_date"`
	EndDate       *time.Time     `gorm:"column:end_date;type:timestamp without time zone" json:"end_date"`
	IsCurrent     *bool          `gorm:"column:is_current;type:boolean;not null;default:true" json:"is_current"`
	VendorID      *string        `gorm:"column:vendor_id;type:uuid" json:"vendor_id"`
	ContactID     *string        `gorm:"column:contact_id;type:uuid" json:"contact_id"`
	ChangeReason  *string        `gorm:"column:change_reason;type:price_change_reason;not null;default:initial" json:"change_reason"`
	Notes         *string        `gorm:"column:notes;type:text" json:"notes"`
	CreatedAt     *time.Time     `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp without time zone" json:"deleted_at"`
	CreatedBy     *string        `gorm:"column:created_by;type:uuid" json:"created_by"`
	UpdatedBy     *string        `gorm:"column:updated_by;type:uuid" json:"updated_by"`
	DeletedBy     *string        `gorm:"column:deleted_by;type:uuid" json:"deleted_by"`
}

// TableName PriceHistory's table name
func (*PriceHistory) TableName() string {
	return TableNamePriceHistory
}
