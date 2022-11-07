package product

import (
	"encoding/csv"
	"linxdatacenter/internal/model"
	"os"
	"strconv"
	"strings"
)

func (s *Storage) GetCSV() (model.ProductList, error) {
	f, err := os.Open(s.file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	products := make(model.ProductList, 0, len(records))
	for _, record := range records {
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}

		rating, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			continue
		}

		products = append(products, model.Product{
			Name:   strings.TrimSpace(record[0]),
			Price:  price,
			Rating: rating,
		})
	}

	return products, nil
}
