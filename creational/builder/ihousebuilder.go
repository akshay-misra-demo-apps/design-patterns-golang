package builder

type IHouseBuilder interface {
	AddBedroom(length, width int, floorType FloorCeramicType) IHouseBuilder
	AddBathroom(length, width int, floorType FloorCeramicType) IHouseBuilder
	AddKitchen(length, width int, floorType FloorCeramicType) IHouseBuilder
	Build() *House
}

func GetHouseBuilder(houseType string) IHouseBuilder {

	if houseType == "standard" {
		return StandardHouseBuilder()
	}

	return StandardHouseBuilder()
}
