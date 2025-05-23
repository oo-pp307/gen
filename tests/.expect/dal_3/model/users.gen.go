// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.
// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.
// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"-"`
	Name         *string    `gorm:"column:name;index:idx_name,priority:1;index:idx_name_company_id,priority:1" json:"-"` // oneline
	Address      *string    `gorm:"column:address" json:"-"`
	RegisterTime *time.Time `gorm:"column:register_time" json:"-"`
	/*
		multiline
		line1
		line2
	*/
	Alive      *bool   `gorm:"column:alive" json:"-"`
	CompanyID  *int64  `gorm:"column:company_id;index:idx_name_company_id,priority:2;default:666" json:"-"`
	PrivateURL *string `gorm:"column:private_url;default:https://a.b.c" json:"-"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
