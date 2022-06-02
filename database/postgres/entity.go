package postgres

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Name     string
}

type Cart struct {
	gorm.Model
	User_id     int
	Status      int
	total_price float64
}

type CartItem struct {
	gorm.Model
	Cart_id    int
	Product_id int
	Quantity   int
}

type Product struct {
	gorm.Model
	Name     string
	Category string
	Price    float64
}

type Transaction struct {
	gorm.Model
	User_id     int
	Cart_id     int
	Status      int
	Total_price float64
}
