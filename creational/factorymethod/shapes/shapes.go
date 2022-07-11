package shapes

import "fmt"

/**
 * Common interface for all shapes.
 */
type Shape interface {
	Draw()
}

type Circle struct {
}

func (c Circle) Draw() {
	fmt.Println("Drawing a circle!")
}

type Square struct {
}

func (s Square) Draw() {
	fmt.Println("Drawing a square!")
}

type Rectangle struct {
}

func (r Rectangle) Draw() {
	fmt.Println("Drawing a rectangle!")
}
