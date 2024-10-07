package factory

import "fmt"

type Product interface {
	Info() string
}

// Concrete products
type product1 struct{}

func (p1 *product1) Info() string {
	return "I am product 1"
}

type product2 struct{}

func (p2 *product2) Info() string {
	return "I am product 2"
}

type product3 struct{}

func (p3 *product3) Info() string {
	return "I am product 3"
}

// Abstract factory
type abstractFactory struct {
}

func (af *abstractFactory) CreateProduct() (Product, error) {
	return nil, fmt.Errorf("this is abstract, needs to be overridden!")
}

// Factory for Product1
type FactoryProduct1 struct{ abstractFactory }

func (fp1 *FactoryProduct1) CreateProduct() (Product, error) {
	return &product1{}, nil
}

// Factory for Product2
type FactoryProduct2 struct{}

func (fp2 *FactoryProduct2) CreateProduct() (Product, error) {
	return &product2{}, nil
}

// Factory for Product3
type FactoryProduct3 struct{}

func (fp3 *FactoryProduct3) CreateProduct() (Product, error) {
	return &product3{}, nil
}
