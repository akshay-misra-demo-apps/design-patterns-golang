package main

import (
	"fmt"
	"reflect"

	"dp.com/design-patterns/creational/abstractfactory"
	"dp.com/design-patterns/creational/builder"
	"dp.com/design-patterns/creational/builder/directors"
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
	fmt.Println("")
	fmt.Println("")

	fmt.Println("builder factory start...")

	fmt.Println("GetHouseBuilder...", reflect.TypeOf(builder.GetHouseBuilder("")))

	// here we have more controlover size and floor type of rooms.
	house := builder.GetHouseBuilder("standard").
		AddBedroom(10, 14, builder.GRANITE).
		AddBedroom(14, 14, builder.MARBLE).
		AddBathroom(7, 7, builder.MOSAIC).
		AddKitchen(10, 10, builder.MARBLE).
		Build()

	fmt.Println("Builder, house: ", house)
	fmt.Println("")
	fmt.Println("")

	house = directors.OneBedroomHouseBuildDirector(builder.GetHouseBuilder("")).BuildHouse()
	fmt.Println("OneBedroomHouseBuildDirector, house: ", house)
	fmt.Println("")
	fmt.Println("")

	house = directors.TwoBedroomHouseBuildDirector(builder.GetHouseBuilder("")).BuildHouse()
	fmt.Println("TwoBedroomHouseBuildDirector, house: ", house)

	fmt.Println("builder factory ends...")

}
