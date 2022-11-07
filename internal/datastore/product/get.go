package product

import (
	"errors"
	"linxdatacenter/internal/model"
	"path/filepath"
)

func (s *Storage) Get() (model.ProductList, error) {
	ext := filepath.Ext(s.file)

	switch ext {
	case ".csv":
		return s.GetCSV()
	case ".json":
		return s.GetJson()
	}

	return nil, errors.New("format not supported")
}
