package domain

type brand struct {
	name           string
	usualGearPower string
}

var (
	Zink              = brand{"アイロニック", "スーパージャンプ短縮"}
	Annaki            = brand{"アナアキ", "インク効率アップ(サブ)"}
	Tentatek          = brand{"アロメ", "インク回復力アップ"}
	Zekko             = brand{"エゾッコ", "スペシャル減少量ダウン"}
	Zekkoree          = brand{"エゾッコリー", "スペシャル減少量ダウン"}
	Emperry           = brand{"エンペリー", "サブ性能UP"}
	Krakon            = brand{"クラーゲス", "イカ移動速度アップ"}
	Inkline           = brand{"シグレニ", "サブ影響軽減"}
	Emberz            = brand{"シチリン", "アクション強化"}
	SprashMob         = brand{"ジモン", "インク効率アップ(メイン)"}
	ToniKensa         = brand{"タタキケンサキ", "インク効率アップ(メイン)"}
	SquidForce        = brand{"バトロイカ", "相手インク影響軽減"}
	Barazushi         = brand{"バラズシ", "アクション強化"}
	Forge             = brand{"フォーリマ", "スペシャル性能アップ"}
	Skalop            = brand{"ホタックス", "復活時間短縮"}
	Firefin           = brand{"ホッコリー", "インク効率アップ(サブ)"}
	Takoroka          = brand{"ヤコ", "スペシャル増加量アップ"}
	Rockenberg        = brand{"ロッケンベルグ", "ヒト移動速度アップ"}
	Cuttlegear        = brand{"アタリメイド", "なし"}
	GrizzcoIndustries = brand{"クマサン商会", "なし"}
	Amiibo            = brand{"amiibo", "non"}
)

func (b *brand) Name() string {
	return b.name
}

func (b *brand) UsualGearPower() string {
	return b.usualGearPower
}

func GetBrandNameList() []string {
	l := make([]string, 0)
	l = append(l, Zink.Name())
	l = append(l, Annaki.Name())
	l = append(l, Tentatek.Name())
	l = append(l, Zekko.Name())
	l = append(l, Zekkoree.Name())
	l = append(l, Emperry.Name())
	l = append(l, Krakon.Name())
	l = append(l, Inkline.Name())
	l = append(l, Emberz.Name())
	l = append(l, SprashMob.Name())
	l = append(l, ToniKensa.Name())
	l = append(l, SquidForce.Name())
	l = append(l, Barazushi.Name())
	l = append(l, Forge.Name())
	l = append(l, Skalop.Name())
	l = append(l, Firefin.Name())
	l = append(l, Takoroka.Name())
	l = append(l, Rockenberg.Name())
	l = append(l, Cuttlegear.Name())
	l = append(l, GrizzcoIndustries.Name())
	l = append(l, Amiibo.Name())
	return l
}
