package domain

type Gesotown struct {
	PickupBrand struct {
		Brand struct {
			Name           string `json:"name"`
			UsualGearPower struct {
				Name string `json:"name"`
			} `json:"usualGearPower"`
		} `json:"brand"`
		SaleEndTime string `json:"saleEndTime"`
	} `json:"pickupBrand"`
}
