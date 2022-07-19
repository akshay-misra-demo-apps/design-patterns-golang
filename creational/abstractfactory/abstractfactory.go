package abstractfactory

import (
	"fmt"

	"dp.com/design-patterns/creational/abstractfactory/shirt"
	"dp.com/design-patterns/creational/abstractfactory/shoe"
)

// Abstract Factory
type ISportsFactory interface {
	MakeShoe(size int) shoe.IShoe
	MakeShirt(size int) shirt.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {

	var factory ISportsFactory

	if brand == "adidas" {
		factory = new(Adidas)
	} else if brand == "nike" {
		factory = new(Nike)
	} else {
		return nil, fmt.Errorf("wrong brand type passed")
	}

	return factory, nil
}
