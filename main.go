package main

import (
	"fmt"

	"dp.com/design-patterns/creational/factorymethod"
)

func main() {
	fmt.Println("starting design patterns...")
	shape := factorymethod.GetShape("circle")
	shape.Draw()
	shape = factorymethod.GetShape("square")
	shape.Draw()
}
