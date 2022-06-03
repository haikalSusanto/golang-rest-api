package cart

import "gorm.io/gorm"

type AddItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	gorm.Model
	User_id     int
	Status      int
	Total_price float64
}

type CartItem struct {
	gorm.Model
	Cart_id    int
	Product_id int
	Quantity   int
}

type User struct {
	ID       int
	username string
}

type OngoingCart struct {
	ID     int
	Status string
}

// type CartItem struct {
// 	ID         int
// 	Cart_id    int
// 	Product_id int
// 	Quantity   int
// }
