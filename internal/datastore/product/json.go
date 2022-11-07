package product

import (
	"encoding/json"
	"io"
	"linxdatacenter/internal/model"
	"os"
)

func (s *Storage) GetJson() (model.ProductList, error) {
	f, err := os.Open(s.file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	byteValue, _ := io.ReadAll(f)

	var products []model.Product

	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
