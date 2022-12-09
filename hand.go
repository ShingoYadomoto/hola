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

func (hc HupaiCalculater) menzen() AllHands {
	if hc.standard.IsMenzen() {
		return AllHands{門前清自摸和}
	}
	return AllHands{}
}

func (hc HupaiCalculater) zhuangfengpai() AllHands {
	if hc.standard.IsZhuangfengpai(hc.zhuangfeng) {
		return AllHands{翻牌場風}
	}
	return AllHands{}
}

func (hc HupaiCalculater) zifengpai() AllHands {
	if hc.standard.IsZifeng(hc.zifeng) {
		return AllHands{翻牌自風}
	}
	return AllHands{}
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

		all = append(all, calculater.menzen()...)        // 門前
		all = append(all, calculater.zhuangfengpai()...) // 場風
		all = append(all, calculater.zifengpai()...)     // 自風
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
