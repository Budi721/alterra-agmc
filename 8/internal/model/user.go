package model

// User representation entities
type User struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Email    string
	Password string
}
