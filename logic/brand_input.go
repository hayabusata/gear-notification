package logic

import "gear-notification/domain"

func OutputBrandFromNumber(numbers []int) []any {
	NumberAndBrandMap := map[int]any{
		1:  domain.Zink,
		2:  domain.Annaki,
		3:  domain.Tentatek,
		4:  domain.Zekko,
		5:  domain.Zekkoree,
		6:  domain.Emperry,
		7:  domain.Krakon,
		8:  domain.Inkline,
		9:  domain.Emberz,
		10: domain.SprashMob,
		11: domain.ToniKensa,
		12: domain.SquidForce,
		13: domain.Barazushi,
		14: domain.Forge,
		15: domain.Skalop,
		16: domain.Firefin,
		17: domain.Takoroka,
		18: domain.Rockenberg,
	}

	brandList := make([]any, 0)
	for _, i := range numbers {
		brandList = append(brandList, NumberAndBrandMap[i])
	}

	return brandList
}
