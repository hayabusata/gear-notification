package main

import (
	"fmt"
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

	lc := notification.NewLineClient("https://notify-api.line.me/api/notify")

	// input
	brandInputs := []string{"ロッケンベルグ", "ないん"}
	badInput, result := logic.ConfirmBrand(brandInputs)
	if !result {
		if err = lc.PostBadBrandInputMessage(badInput); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(err)
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
		// notification.PostGearMessage(gear)
		fmt.Println(gear)
	}
	// notification.PostPickupBrandMessage(*geso)
}
