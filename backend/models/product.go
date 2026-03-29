package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Code      string         `gorm:"uniqueIndex;size:50;not null" json:"code"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Spec      string         `gorm:"size:100" json:"spec"`
	Unit      string         `gorm:"size:20" json:"unit"`
	Category  string         `gorm:"size:50" json:"category"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type ProductCreate struct {
	Code     string `json:"code" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Spec     string `json:"spec"`
	Unit     string `json:"unit"`
	Category string `json:"category"`
}

type ProductUpdate struct {
	Name     string `json:"name"`
	Spec     string `json:"spec"`
	Unit     string `json:"unit"`
	Category string `json:"category"`
}
