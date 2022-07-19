package abstractfactory

import (
	"dp.com/design-patterns/creational/abstractfactory/shirt"
	"dp.com/design-patterns/creational/abstractfactory/shoe"
)

// Nike Factory
type Nike struct{}

func (n Nike) MakeShoe(size int) shoe.IShoe {

	shoe := new(shoe.NikeShoe)
	shoe.SetLogo("Nike")
	shoe.SetSize(size)

	return shoe

}

func (n Nike) MakeShirt(size int) shirt.IShirt {

	shirt := new(shirt.NikeShirt)
	shirt.SetLogo("Nike")
	shirt.SetSize(size)

	return shirt
}
