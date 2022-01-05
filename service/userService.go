package service

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"poc-model/model"
	"time"
)

type UserService interface {
	GetUserById(userId uint) model.User
	Update(user model.User) model.User
	AddUser(userReq model.User) (*model.User, error)
	UpdateUser(id string, userReq model.User) (*model.User, error)
	DeleteUser(id string) error
	GetUser(userId uint) model.User
	GetUserByEmailAndType(email string, userType string) (*model.User, error)
	GetIngredientById(id uint) (*model.Ingredient, error)
}
type userInformation struct {
	DB *gorm.DB
}

func UserServiceInitialize(db *gorm.DB) UserService {
	return &userInformation{
		DB: db,
	}
}
func (info *userInformation) AddUser(userReq model.User) (*model.User, error) {
	if err := info.DB.Save(&userReq).Error; err != nil {
		return &userReq, err
	}

	return &userReq, nil
}

func (info *userInformation) UpdateUser(id string, userReq model.User) (*model.User, error) {
	if err := info.DB.Find(&userReq, id).Error; err != nil {
		return &userReq, err
	}

	if err := info.DB.Save(&userReq).Error; err != nil {

		return &userReq, err
	}

	return &userReq, nil
}

func (info *userInformation) DeleteUser(id string) error {
	user := model.User{}
	if err := info.DB.First(&user, id).Error; err != nil {
		return err
	}

	user.IsActive = false

	if err := info.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userInformation) GetUser(userId uint) model.User {
	var user = model.User{}
	u.DB.Preload(clause.Associations).First(&user,userId)
	fmt.Println(userId)
	fmt.Println(&user)
	return user
}

func (u *userInformation) GetUserById(userId uint) model.User {
	return u.GetUser(userId)
}

func (u *userInformation) Update(user model.User) model.User {
	user.UpdatedAt = time.Now().UTC()
	u.DB.Save(user)
	return user
}

func (u *userInformation) GetUserByEmailAndType(email string, userType string) (*model.User, error) {

	var user = model.User{}
	result := u.DB.Where("email = ? AND type = ?", email, userType).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *userInformation) GetIngredientById(id uint) (*model.Ingredient, error) {

	var ingredient = model.Ingredient{}
	result := u.DB.Where("id = ?", id).First(&ingredient)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ingredient, nil
}