package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CourseStatus string

const (
	DRAFT     CourseStatus = "DRAFT"
	PUBLISHED              = "PUBLISHED"
	ENDED                  = "ENDED"
)

func (e CourseStatus) ToString() string {
	return string(e)
}

func (t CourseStatus) Value() (driver.Value, error) {
	switch t {
	case DRAFT, PUBLISHED, ENDED: //valid case
		return string(t), nil
	}
	return nil, errors.New("Invalid course status value") //else is invalid
}

func (t *CourseStatus) Scan(value interface{}) error {
	var pt CourseStatus
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid data for course status")
	}
	pt = CourseStatus(string(value.([]byte))) //convert type from string to ProductType
	switch pt {
	case DRAFT, PUBLISHED, ENDED: //valid case
		*t = pt
		return nil
	}
	return fmt.Errorf("Invalid course status value :%s", st) //else is invalid
}
// one to many -> CourseMenu
// one to one -> CoverImage
type Course struct {
	gorm.Model
	ChefID                   uint
	Chef                     Chef    `json:"chef"`
	Location                 *ChefLocation  `gorm:"foreignkey:CourseID;references:ID"`
	TypeOfCuisines           *[]TypeOfCuisine `gorm:"many2many:course_typeOfCuisines;"`
	CourseMenus           	 *[]CourseMenu `gorm:"foreignkey:CourseID;references:ID"`
	CoverImage               *Image	`gorm:"foreignkey:CourseID;references:ID"`
	Status                   CourseStatus `gorm:"default:'DRAFT';not null" json:"status"`
}

// master data many to many
type TypeOfCuisine struct {
	gorm.Model
	Name      *string   `gorm:"type:varchar(100) not null" json:"name"`
	IsActive  bool      `gorm:"not null" json:"isActive"`
}

// many to one -> CourseID
type CourseMenu struct {
	gorm.Model
	CourseID 	uint
	Name      string    `gorm:"not null" json:"name"`
	Sequence  int       `gorm:"not null" json:"-"`
}

type Chef struct {
	gorm.Model
	UserID         uint
	Name           string `gorm:"type:varchar(30);not null"`
}

type ChefLocation struct {
	gorm.Model
	CourseID		uint
	ChefID          uint           `gorm:"not null" json:"-"`
	Name            string         `gorm:"size:30,not null" json:"name"`
	Latitude        string         `gorm:"size:64" json:"latitude"`
	Longitude       string         `gorm:"size:64" json:"longitude"`
	Address         string         `gorm:"size:255" json:"address"`
	IsParking       bool           `gorm:"default:false" json:"isParking"`
	ParkingLocation string         `gorm:"size:255" json:"parkingLocation"`
}
