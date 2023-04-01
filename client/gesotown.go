package client

import (
	"encoding/json"
	"gear-notification/domain"
	"io"
	"log"
	"net/http"
)

func GetGesotownGearList() (*domain.Gesotown, error) {
	res, err := http.Get("https://spla3.yuu26.com/api/schedule")
	log.Printf("response log: %v", res)
	if err != nil {
		log.Printf("error log: %v", err)
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
