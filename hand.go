package hola_go

type HupaiCalculater struct {
	standard   StandardHoluPattern
	zhuangfeng Zhuangfeng
	zifeng     Zifeng
}

func NewHupaiCalculater(standard StandardHoluPattern, zhuangfeng Zhuangfeng, zifeng Zifeng) *HupaiCalculater {
	return &HupaiCalculater{
		standard:   standard,
		zhuangfeng: zhuangfeng,
		zifeng:     zifeng,
	}
}

func (hc HupaiCalculater) menzen() HandType {
	if hc.standard.IsMenzen() {
		return 門前清自摸和
	}
	return 0
}

func (hc HupaiCalculater) fengpai() []HandType {
	hands := []HandType{}
	if hc.standard.IsZhuangfengpai(hc.zhuangfeng) {
		hands = append(hands, 翻牌場風)
	}
	if hc.standard.IsZifeng(hc.zifeng) {
		hands = append(hands, 翻牌自風)
	}
	if hc.standard.HasKotsu(白) {
		hands = append(hands, 翻牌白)
	}
	if hc.standard.HasKotsu(發) {
		hands = append(hands, 翻牌發)
	}
	if hc.standard.HasKotsu(中) {
		hands = append(hands, 翻牌中)
	}
	return hands
}

type FullHupaiCalculater struct {
	fullParrern FullHoluPattern
	zhuangfeng  Zhuangfeng
	zifeng      Zifeng
}

func (fhc FullHupaiCalculater) Hupai() AllHands {
	all := AllHands{}
	all = append(all, fhc.kokushi()...)
	all = append(all, fhc.titoitsu()...)
	for _, standard := range fhc.fullParrern.Standard {
		calculater := NewHupaiCalculater(standard, fhc.zhuangfeng, fhc.zifeng)

		all = append(all, calculater.menzen())     // 門前
		all = append(all, calculater.fengpai()...) // 場風・自風・白・發・中
	}

	return nil
}

func (fhc FullHupaiCalculater) kokushi() AllHands {
	if fhc.fullParrern.Kokushi != nil {
		if fhc.fullParrern.Kokushi.IsDouble() {
			return AllHands{国士無双十三面}
		}
		return AllHands{国士無双}
	}
	return AllHands{}
}

func (fhc FullHupaiCalculater) titoitsu() AllHands {
	if fhc.fullParrern.Titoitsu != nil {
		return AllHands{七対子}
	}
	return AllHands{}
}
