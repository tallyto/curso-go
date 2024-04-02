package entity

import (
	"errors"
	"time"

	"github.com/tallyto/curso-go/7-apis/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIsInvalid  = errors.New("price is invalid")
	ErrInvalidId       = errors.New("invalid id")
)

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price < 0 {
		return ErrPriceIsInvalid
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	return nil
}

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()

	if err != nil {
		return nil, err
	}

	return product, nil
}
