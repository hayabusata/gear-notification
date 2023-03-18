package domain

type Gear struct {
	Id          string
	Price       int
	SaleEndTime string
	Gear        struct {
		Typename         string
		Name             string
		PrimaryGearPower string
		Brand            struct {
			Name string
			Id   string
		}
	}
}
