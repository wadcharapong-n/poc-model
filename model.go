package main

import (
	"gorm.io/gorm"
	"time"
)

type EnvConfig struct {
	Host       string
	Port       string
	DbConfig   MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

type User struct {
	ID                  uint    `json:"id" gorm:"primaryKey"`
	Email               *string `json:"email,omitempty"`
	Password            *string `json:"-"`
	Firstname           *string `json:"firstName,omitempty"`
	Lastname            *string `json:"lastName,omitempty"`
	Label               *string `json:"label,omitempty"`
	ProfileImageID      *uint   `json:"profileImageID,omitempty"`
	NumberOfVaccination *uint8           `json:"numberOfVaccination,omitempty"`
	Type                *string          `json:"type,omitempty"`
	IsActive            bool             `gorm:"default:true" json:"isActive"`
	IsChef              bool             `gorm:"default:false" json:"isChef"`
	CreatedAt           time.Time        `json:"createdAt,omitempty"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
	Allergies           *[]Ingredient `gorm:"polymorphic:Owner;polymorphicValue:master"`
}

type Ingredient struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	OwnerID   int
	OwnerType string
	IsActive bool   `json:"isActive" gorm:"default:true"`
}

type TimeAvailable struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Time      string         `gorm:"size(20),not null" json:"time"`
	Sequence  int            `gorm:"not null" json:"-"`
}

