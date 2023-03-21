package logic

import "gear-notification/domain"

func contains(s string, list []string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}

	return false
}

func FilterGearsByBrand(
	base []domain.Gear,
	brands []string,
) []domain.Gear {
	result := []domain.Gear{}

	for _, gear := range base {
		if contains(gear.Gear.Brand.Name, brands) {
			result = append(result, gear)
		}
	}

	return result
}

func FilterGearsByMainPower(
	base []domain.Gear,
	mainPowers []string,
) []domain.Gear {
	result := []domain.Gear{}

	for _, gear := range base {
		if contains(gear.Gear.Brand.Name, mainPowers) {
			result = append(result, gear)
		}
	}

	return result
}

func FilterGearsByGearName(
	base []domain.Gear,
	gearNames []string,
) []domain.Gear {
	result := []domain.Gear{}

	for _, gear := range base {
		if contains(gear.Gear.Brand.Name, gearNames) {
			result = append(result, gear)
		}
	}

	return result
}
