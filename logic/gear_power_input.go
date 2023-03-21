package logic

import "gear-notification/domain"

func OutputGearPowerFromNumber(numbers []int) []any {
	NumberAndGearPowerMap := map[int]any{
		1:  domain.SuperJumpReduction,
		2:  domain.SubInkEfficiencyUp,
		3:  domain.InkRecovery,
		4:  domain.SpecialDecreaseDown,
		5:  domain.SubPerformanceUp,
		6:  domain.SquidMovementSpeedUp,
		7:  domain.SubImpactReduction,
		8:  domain.ActionEnhancement,
		9:  domain.MainInkEfficiencyUp,
		10: domain.SafetyShoes,
		11: domain.SpecialPerformanceUp,
		12: domain.ResurrectionTimeReduction,
		13: domain.SpecialIncrementUp,
		14: domain.HumanMovementSpeedUp,
	}

	gearPowerList := make([]any, 0)
	for _, i := range numbers {
		gearPowerList = append(gearPowerList, NumberAndGearPowerMap[i])
	}

	return gearPowerList
}
