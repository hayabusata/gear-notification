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
		// BrandGears  struct {} `json:"brandGears"`
	} `json:"pickupBrand"`
	LimitedGears []struct {
		Id          string `json:"id"`
		Price       int    `json:"price"`
		SaleEndTime string `json:"saleEndTime"`
		Gear        struct {
			Typename         string `json:"typename"`
			Name             string `json:"name"`
			PrimaryGearPower string `json:"primaryGearPower"`
			Brand            struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"brand"`
		} `json:"gear"`
	} `json:"limitedGears"`
}
