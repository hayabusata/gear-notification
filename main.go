package main

import (
	"encoding/json"
	"fmt"
	"gear-notification/domain"
	"gear-notification/notification"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.koukun.jp/splatoon/3/geso/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var geso domain.Gesotown
	if err = json.Unmarshal(body, &geso); err != nil {
		panic(err)
	}

	fmt.Println(&geso)

	/* line */
	notification.PostPickupBrandMessage(geso)

	for i := 0; i < 6; i++ {
		gear := domain.Gear(geso.LimitedGears[i])
		notification.PostGearMessage(gear)
	}

}
