package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string
	Category string
	Price    float64
}

type ListProduct struct {
	Products []Product
}
