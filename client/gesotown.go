package client

import (
	"encoding/json"
	"gear-notification/domain"
	"io"
	"net/http"
)

func GetGesotownGearList() (*domain.Gesotown, error) {
	res, err := http.Get("https://api.koukun.jp/splatoon/3/geso/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var geso domain.Gesotown
	if err = json.Unmarshal(body, &geso); err != nil {
		return nil, err
	}

	return &geso, nil
}
