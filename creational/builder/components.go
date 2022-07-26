package builder

import "fmt"

type Room struct {
	length    int
	width     int
	floorType FloorCeramicType
}

type FloorCeramicType int

const (
	MARBLE FloorCeramicType = iota
	MOSAIC
	GRANITE
)

type Bedroom struct {
	Room
}

type Bathroom struct {
	Room
}

type Kitchen struct {
	Room
}

type House struct { // interface?
	bedroom  []*Bedroom
	bathroom []*Bathroom
	kitchen  []*Kitchen
}

func (h *House) addBedroom(length, width int, floorType FloorCeramicType) {
	b := new(Bedroom)
	b.length = length
	b.width = width
	b.floorType = floorType

	fmt.Println("addBedroom, house: ", h)

	if h.bedroom == nil {
		h.bedroom = make([]*Bedroom, 0)
	}

	h.bedroom = append(h.bedroom, b)
}

func (h *House) addBathroom(length, width int, floorType FloorCeramicType) {
	b := new(Bathroom)
	b.length = length
	b.width = width
	b.floorType = floorType

	if h.bathroom == nil {
		h.bathroom = make([]*Bathroom, 0)
	}

	h.bathroom = append(h.bathroom, b)
}

func (h *House) addKitchen(length, width int, floorType FloorCeramicType) {
	k := new(Kitchen)
	k.length = length
	k.width = width
	k.floorType = floorType

	if h.kitchen == nil {
		h.kitchen = make([]*Kitchen, 0)
	}

	h.kitchen = append(h.kitchen, k)
}
