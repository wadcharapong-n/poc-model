package model

import (
	"gorm.io/gorm"
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
//case one to one  ->  ProfileImage
//case many to many -> Allergies
type User struct {
	gorm.Model
	Email               *string `json:"email,omitempty"`
	Password            *string `json:"-"`
	Firstname           *string `json:"firstName,omitempty"`
	Lastname            *string `json:"lastName,omitempty"`
	Label               *string `json:"label,omitempty"`
	ProfileImage      	*Image	`gorm:"foreignkey:UserId;references:ID"`
	NumberOfVaccination *uint8           `json:"numberOfVaccination,omitempty"`
	IsChef              bool             `gorm:"default:false" json:"isChef"`
	IsActive 			bool  `gorm:"default:false" json:"isActive"`
	Allergies           *[]Ingredient `gorm:"many2many:user_allergies;"`
	Chef		*Chef
}

type Ingredient struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
}


type Image struct {
	gorm.Model
	CourseID 		*uint
	UserId			*uint
	ImageUrl     	string     `gorm:"type:varchar(255) not null" json:"imageUrl"`
}
