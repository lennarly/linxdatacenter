package model

type (
	ProductList []Product

	Product struct {
		Name   string  `json:"Product"` // Наименование продукта.
		Price  float64 `json:"Price"`   // Цена продукта.
		Rating float64 `json:"Rating"`  // Рейтинг продукта.
	}
)

func (p *ProductList) GetHighPriced() *Product {
	if len(*p) == 0 {
		return &Product{}
	}

	bestProduct := (*p)[0]

	for _, product := range *p {
		if product.Price > bestProduct.Price {
			bestProduct = product
		}
	}

	return &bestProduct
}

func (p *ProductList) GetHighRanked() *Product {
	if len(*p) == 0 {
		return &Product{}
	}

	bestProduct := (*p)[0]

	for _, product := range *p {
		if product.Rating > bestProduct.Rating {
			bestProduct = product
		}
	}

	return &bestProduct
}
