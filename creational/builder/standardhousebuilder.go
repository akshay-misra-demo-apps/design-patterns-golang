package builder

import "fmt"

type standardHouseBuilder struct {
	house *House
}

func StandardHouseBuilder() *standardHouseBuilder {
	return &standardHouseBuilder{
		house: new(House),
	}
}

func (b *standardHouseBuilder) AddBedroom(length, width int, floorType FloorCeramicType) IHouseBuilder {
	fmt.Println("AddBedroom, house: ", b.house)
	b.house.addBedroom(length, width, floorType)

	return b
}

func (b *standardHouseBuilder) AddBathroom(length, width int, floorType FloorCeramicType) IHouseBuilder {
	b.house.addBathroom(length, width, floorType)

	return b
}

func (b *standardHouseBuilder) AddKitchen(length, width int, floorType FloorCeramicType) IHouseBuilder {
	b.house.addKitchen(length, width, floorType)

	return b
}

func (b *standardHouseBuilder) Build() *House {
	return b.house
}
