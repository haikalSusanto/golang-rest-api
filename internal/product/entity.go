package product

type Product struct {
	ID       string
	Name     string
	Category string
	Price    float64
}

type ListProduct struct {
	Products []Product
}
