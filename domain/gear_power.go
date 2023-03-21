package domain

type gearPower struct {
	name               string
	easyPowerBrandName []string
}

var (
	SuperJumpReduction        = gearPower{"スーパージャンプ短縮", []string{"アイロニック"}}
	SubInkEfficiencyUp        = gearPower{"インク効率アップ(サブ)", []string{"アナアキ", "ホッコリー"}}
	InkRecovery               = gearPower{"インク回復力アップ", []string{"アロメ"}}
	SpecialDecreaseDown       = gearPower{"スペシャル減少量ダウン", []string{"エゾッコ", "エゾッコリー"}}
	SubPerformanceUp          = gearPower{"サブ性能UP", []string{"エンペリー"}}
	SquidMovementSpeedUp      = gearPower{"イカ移動速度アップ", []string{"クラーゲス"}}
	SubImpactReduction        = gearPower{"サブ影響軽減", []string{"シグレニ"}}
	ActionEnhancement         = gearPower{"アクション強化", []string{"シチリン", "バラズシ"}}
	MainInkEfficiencyUp       = gearPower{"インク効率アップ(メイン)", []string{"ジモン", "タタキケンサキ"}}
	SafetyShoes               = gearPower{"相手インク影響軽減", []string{"バトロイカ"}}
	SpecialPerformanceUp      = gearPower{"スペシャル性能アップ", []string{"フォーリマ"}}
	ResurrectionTimeReduction = gearPower{"復活時間短縮", []string{"ホタックス"}}
	SpecialIncrementUp        = gearPower{"スペシャル増加量アップ", []string{"ヤコ"}}
	HumanMovementSpeedUp      = gearPower{"ヒト移動速度アップ", []string{"ロッケンベルグ"}}
)

func (gp *gearPower) Name() string {
	return gp.name
}

func (gp *gearPower) EasyPowerBrandName() []string {
	return gp.easyPowerBrandName
}

func GetGearPowerNameList() []string {
	l := make([]string, 0)
	l = append(l, SuperJumpReduction.Name())
	l = append(l, SubInkEfficiencyUp.Name())
	l = append(l, InkRecovery.Name())
	l = append(l, SpecialDecreaseDown.Name())
	l = append(l, SubPerformanceUp.Name())
	l = append(l, SquidMovementSpeedUp.Name())
	l = append(l, SubImpactReduction.Name())
	l = append(l, ActionEnhancement.Name())
	l = append(l, MainInkEfficiencyUp.Name())
	l = append(l, SafetyShoes.Name())
	l = append(l, SpecialPerformanceUp.Name())
	l = append(l, ResurrectionTimeReduction.Name())
	l = append(l, SpecialIncrementUp.Name())
	l = append(l, HumanMovementSpeedUp.Name())
	return l
}
