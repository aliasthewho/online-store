package products

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (product *Product) GetInfo() string {
	return product.Name
}
