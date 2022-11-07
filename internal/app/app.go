package app

import (
	"fmt"
	"linxdatacenter/internal/datastore/product"
)

func Run() {
	file := product.New("internal/testdata/red.json")

	products, err := file.Get()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	fmt.Printf("The highest priced product: %v\n", products.GetHighPriced())
	fmt.Printf("Top rated product: %v\n", products.GetHighRanked())
}
