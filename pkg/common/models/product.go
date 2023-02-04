package models

type Product struct {
	Id    int     `json:"id" gorm:"primary_key"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int32   `json:"stock"`
}
