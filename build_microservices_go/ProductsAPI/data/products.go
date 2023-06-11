package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

// returns the list of products
func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Coffee Granita",
		Description: "Practically an Icecream Coffee",
		Price:       1.95,
		SKU:         "70F",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Hyper Energy booster",
		Price:       0.95,
		SKU:         "70A",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
