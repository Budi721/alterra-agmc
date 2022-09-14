package models

type Book struct {
	ID     uint   `json:"id" gorm:"primarykey"`
	Title  string `json:"title,omitempty" form:"title"`
	Author string `json:"author,omitempty" form:"author"`
	Price  uint   `json:"price,omitempty" form:"price"`
}
