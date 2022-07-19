package abstractfactory

import (
	"fmt"

	"dp.com/design-patterns/creational/abstractfactory/shirt"
	"dp.com/design-patterns/creational/abstractfactory/shoe"
)

// Adidas Factory
type Adidas struct{}

func (a Adidas) MakeShoe(size int) shoe.IShoe {
	fmt.Println("Adidas factory, MakeShoe, size: ", size)

	shoe := new(shoe.AdidasShoe)

	shoe.SetLogo("Adidas")
	shoe.SetSize(size)

	return shoe

}

func (a Adidas) MakeShirt(size int) shirt.IShirt {
	fmt.Println("Adidas factory, MakeShirt, size: ", size)

	shirt := new(shirt.AdidasShirt)
	shirt.SetLogo("Adidas")
	shirt.SetSize(size)

	return shirt
}
