package dto

import (
	"encoding/json"
	"io"
)

// createProductRequest é uma representação do request body para criação de um Novo Produto
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	createProductRequest := CreateProductRequest{}
	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}
	return &createProductRequest, nil
}
