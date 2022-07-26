package directors

import "dp.com/design-patterns/creational/builder"

type twoBedroomHouseBuildDirector struct {
	b builder.IHouseBuilder
}

func TwoBedroomHouseBuildDirector(builder builder.IHouseBuilder) *twoBedroomHouseBuildDirector {
	return &twoBedroomHouseBuildDirector{
		b: builder,
	}
}

func (d *twoBedroomHouseBuildDirector) BuildHouse() *builder.House {
	return d.b.AddBathroom(12, 16, builder.GRANITE).
		AddBedroom(12, 16, builder.MARBLE).
		AddBedroom(10, 12, builder.MARBLE). // second bedroom not getting appended in house.
		AddKitchen(8, 10, builder.MOSAIC).
		Build()
}
