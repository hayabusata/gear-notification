package main

import (
	"gear-notification/client"
	"gear-notification/domain"
	"gear-notification/logic"
	"gear-notification/notification"
)

func main() {
	geso, err := client.GetGesotownGearList()
	if err != nil {
		panic(err)
	}

	// filter
	brands := []string{"ロッケンベルグ"}
	gears := []domain.Gear{}
	for _, gear := range geso.LimitedGears {
		gears = append(gears, domain.Gear(gear))
	}

	brandFilteredGears := logic.FilterGearsByBrand(gears, brands)

	// notify
	for _, gear := range brandFilteredGears {
		notification.PostGearMessage(gear)
	}
	// notification.PostPickupBrandMessage(*geso)
}
