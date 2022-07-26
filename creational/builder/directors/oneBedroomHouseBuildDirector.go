package directors

import "dp.com/design-patterns/creational/builder"

type oneBedroomHouseBuildDirector struct {
	b builder.IHouseBuilder
}

func OneBedroomHouseBuildDirector(builder builder.IHouseBuilder) *oneBedroomHouseBuildDirector {
	return &oneBedroomHouseBuildDirector{
		b: builder,
	}
}

func (d *oneBedroomHouseBuildDirector) BuildHouse() *builder.House {
	return d.b.AddBathroom(12, 16, builder.GRANITE).
		AddBedroom(12, 16, builder.MARBLE).
		AddKitchen(8, 10, builder.MOSAIC).
		Build()
}
