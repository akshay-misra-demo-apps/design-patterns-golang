package factorymethod

import "dp.com/design-patterns/creational/factorymethod/shapes"

func GetShape(s string) shapes.Shape {

	var shape shapes.Shape

	switch s {
	case "circle":
		shape = new(shapes.Circle)
	case "square":
		shape = new(shapes.Square)
	case "rectangle":
		shape = new(shapes.Rectangle)
	}

	return shape
}
