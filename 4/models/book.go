package models

// Book representation entities for book endpoint
type Book struct {
	ID     uint   `json:"id" validate:"required"`
	Title  string `json:"title,omitempty" form:"title" validate:"required"`
	Author string `json:"author,omitempty" form:"author" validate:"required"`
	Price  uint   `json:"price,omitempty" form:"price" validate:"required"`
}
