package products

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Static data for three products
var Products = []Product{
	{ID: 1, Name: "Apple Watch Series 10", Price: 449.99},
	{ID: 2, Name: "iPad Air ", Price: 699.99},
	{ID: 3, Name: "Airpods Pro 2", Price: 279.99},
}

func FindProductByID(products []Product, id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
