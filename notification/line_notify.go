package notification

import (
	"gear-notification/domain"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type LineClient struct {
	accessToken string
	url         string
}

func NewLineClient(u string) *LineClient {
	lc := new(LineClient)
	lc.accessToken = os.Getenv("line-access-token")
	lc.url = u
	return lc
}

func (lc *LineClient) PostBadBrandInputMessage(brandName string) error {
	msg := `\n"` + brandName + `"は存在しないブランドです`

	u, err := url.ParseRequestURI(lc.url)
	if err != nil {
		return err
	}
	c := &http.Client{}
	form := url.Values{}
	form.Add("message", msg)

	msgBody := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", u.String(), msgBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+lc.accessToken)

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (lc *LineClient) PostPickupBrandMessage(geso domain.Gesotown) error {
	msg := "\n本日のピックアップブランド\n" +
		geso.PickupBrand.Brand.Name + "\n" +
		"(" + geso.PickupBrand.Brand.UsualGearPower.Name + ")\n" +
		geso.PickupBrand.SaleEndTime + "まで\n" +
		"https://api.lp1.av5ja.srv.nintendo.net/gesotown"

	accessToken := os.Getenv("line-access-token")
	URL := "https://notify-api.line.me/api/notify"

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}
	c := &http.Client{}
	form := url.Values{}
	form.Add("message", msg)

	msgBody := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", u.String(), msgBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (lc *LineClient) PostGearMessage(gear domain.Gear) error {
	msg := "\n本日のギア\n" +
		gear.Gear.Name + "[" + gear.Gear.Brand.Name + "]" + "\n" +
		"(" + gear.Gear.PrimaryGearPower + ")\n" +
		gear.SaleEndTime + "まで\n" +
		"https://api.lp1.av5ja.srv.nintendo.net/gesotown"

	u, err := url.ParseRequestURI(lc.url)
	if err != nil {
		return err
	}
	c := &http.Client{}
	form := url.Values{}
	form.Add("message", msg)

	msgBody := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", u.String(), msgBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+lc.accessToken)

	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}
