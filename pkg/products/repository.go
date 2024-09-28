package products

type ProductRepository interface {
	GetAll() []Product
	GetByID(id int) *Product
}

type InMemoryProductRepository struct {
	products []Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: []Product{
			{ID: 1, Name: "Laptop", Price: 1500},
			{ID: 2, Name: "Smartphone", Price: 800},
		},
	}
}

func (r *InMemoryProductRepository) GetAll() []Product {
	return r.products
}

func (r *InMemoryProductRepository) GetByID(id int) *Product {
	for _, p := range r.products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
