package notification

import (
	"gear-notification/domain"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PostMessage(geso domain.Gesotown) {
	msg := "\n本日のピックアップブランド\n" +
		geso.PickupBrand.Brand.Name + "\n" +
		"(" + geso.PickupBrand.Brand.UsualGearPower.Name + ")\n" +
		geso.PickupBrand.SaleEndTime + "まで"

	accessToken := os.Getenv("line-access-token")
	URL := "https://notify-api.line.me/api/notify"

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		log.Fatal(err)
	}
	c := &http.Client{}
	form := url.Values{}
	form.Add("message", msg)

	msgBody := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", u.String(), msgBody)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
