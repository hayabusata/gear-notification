package main

import (
	"gear-notification/client"
	"gear-notification/domain"
	"gear-notification/logic"
	"gear-notification/notification"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	geso, err := client.GetGesotownGearList()
	if err != nil {
		panic(err)
	}

	lc := notification.NewLineClient("https://notify-api.line.me/api/notify")

	// input
	// brandInputs := []string{"ロッケンベルグ", "ないん"}
	// badInput, result := logic.ValidateBrand(brandInputs)
	// if !result {
	// 	if err = lc.PostBadBrandInputMessage(badInput); err != nil {
	// 		panic(err)
	// 	}
	// }

	// filter
	brands := []string{"ヤコ"}

	gears := []domain.Gear{}
	for _, gear := range geso.LimitedGears {
		gears = append(gears, domain.Gear(gear))
	}

	brandFilteredGears := logic.FilterGearsByBrand(gears, brands)

	// line bot notify
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, gear := range brandFilteredGears {
		content := "本日のギア\n" +
			gear.Gear.Name + "[" + gear.Gear.Brand.Name + "]" + "\n" +
			"(" + gear.Gear.PrimaryGearPower + ")\n" +
			gear.SaleEndTime + "まで\n"
		msg := linebot.NewTextMessage(content)

		if _, err := bot.BroadcastMessage(msg).Do(); err != nil {
			log.Fatal(err)
		}
	}

	// line notify
	// for _, gear := range brandFilteredGears {
	// 	if err = lc.PostGearMessage(gear); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(gear)
	// }
}
