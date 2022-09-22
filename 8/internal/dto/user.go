package dto

// User representation entities for user endpoint
type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email"  validate:"required,email"`
	Password string `json:"password" form:"password"  validate:"required"`
}

// UserUpdate representation dto for update endpoint
type UserUpdate struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// UserLogin representation dto for login endpoint
type UserLogin struct {
	Email    string `json:"email" form:"email"  validate:"required,email"`
	Password string `json:"password" form:"password"  validate:"required"`
}
