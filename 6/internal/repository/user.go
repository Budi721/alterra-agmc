package repository

import (
	"github.com/Budi721/alterra-agmc/v6/internal/model"
	"gorm.io/gorm"
)

type User interface {
	LoginUser(email string, password string) (*model.User, error)
	GetUsers() ([]model.User, error)
	GetUser(id uint) (*model.User, error)
	CreateUser(name string, email string, password string) (*model.User, error)
	UpdateUser(id uint, name string, email string, password string) (*model.User, error)
	DeleteUser(id uint) (*model.User, error)
}

type user struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &user{Db: db}
}

func (u *user) LoginUser(email string, password string) (*model.User, error) {
	var user model.User

	if err := u.Db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func (u *user) GetUsers() ([]model.User, error) {
	var users []model.User

	if err := u.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return users, nil
}

func (u *user) GetUser(id uint) (*model.User, error) {
	var user model.User

	if err := u.Db.First(&user, id).Error; err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func (u *user) CreateUser(name string, email string, password string) (*model.User, error) {
	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := u.Db.Create(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func (u *user) UpdateUser(id uint, name string, email string, password string) (*model.User, error) {
	user := model.User{
		ID: id,
	}
	if err := u.Db.First(&user).Error; err != nil {
		return &model.User{}, err
	}

	user.Name = name
	user.Email = email
	user.Password = password

	if err := u.Db.Save(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func (u *user) DeleteUser(id uint) (*model.User, error) {
	user := model.User{
		ID: id,
	}

	if err := u.Db.First(&user).Error; err != nil {
		return &model.User{}, err
	}

	if err := u.Db.Delete(&model.User{}, id).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}
