package main

import (
	"fmt"

	"dp.com/design-patterns/creational/abstractfactory"
	"dp.com/design-patterns/creational/factorymethod"
)

func main() {
	fmt.Println("starting design patterns...")

	fmt.Println("factory method start...")
	shape := factorymethod.GetShape("circle")
	shape.Draw()
	shape = factorymethod.GetShape("square")
	shape.Draw()
	fmt.Println("factory method end...")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("abstract factory start...")
	factory, err := abstractfactory.GetSportsFactory("adidas")

	if err != nil {
		fmt.Println("abstract factory error: ", err.Error())
	}

	if factory != nil {

		shirt := factory.MakeShirt(38)
		shoe := factory.MakeShoe(8)
		fmt.Println("Shirt: ", shirt, ", Shoe: ", shoe)
		fmt.Println("")
	}

	factory, err = abstractfactory.GetSportsFactory("nike")
	if err != nil {
		fmt.Println("abstract factory error: ", err.Error())
	}

	if factory != nil {

		shirt := factory.MakeShirt(38)
		shoe := factory.MakeShoe(8)
		fmt.Println("Shirt: ", shirt, ", Shoe: ", shoe)
		fmt.Println("")
	}

	factory, err = abstractfactory.GetSportsFactory("niki")
	if err != nil {
		fmt.Println("abstract factory error: ", err.Error())
	}

	if factory != nil {
		shirt := factory.MakeShirt(38)
		shoe := factory.MakeShoe(8)
		fmt.Println("Shirt: ", shirt, ", Shoe: ", shoe)
		fmt.Println("")
	}

	fmt.Println("abstract factory ends ...")

}
