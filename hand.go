package hola_go

type HupaiCalculater struct {
	standard   StandardHoluPattern
	zhuangfeng Zhuangfeng
	zifeng     Zifeng
	isTsumo    bool
}

func NewHupaiCalculater(standard StandardHoluPattern, zhuangfeng Zhuangfeng, zifeng Zifeng, isTsumo bool) *HupaiCalculater {
	return &HupaiCalculater{
		standard:   standard,
		zhuangfeng: zhuangfeng,
		zifeng:     zifeng,
		isTsumo:    isTsumo,
	}
}

func (hc HupaiCalculater) tsumo() []HandType {
	if hc.standard.IsMenzen() && hc.isTsumo {
		return []HandType{門前清自摸和}
	}
	return []HandType{}
}

func (hc HupaiCalculater) fengpai() []HandType {
	hands := []HandType{}
	if hc.standard.IsZhuangfengpai(hc.zhuangfeng) {
		hands = append(hands, 翻牌場風)
	}
	if hc.standard.IsZifeng(hc.zifeng) {
		hands = append(hands, 翻牌自風)
	}
	if hc.standard.HasSpecificKotsuOrKantsu(白) {
		hands = append(hands, 翻牌白)
	}
	if hc.standard.HasSpecificKotsuOrKantsu(發) {
		hands = append(hands, 翻牌發)
	}
	if hc.standard.HasSpecificKotsuOrKantsu(中) {
		hands = append(hands, 翻牌中)
	}
	return hands
}

func (hc HupaiCalculater) pinfu() []HandType { panic("not implemented") }

func (hc HupaiCalculater) tanyao() []HandType {
	if !hc.standard.HasYaojiu() {
		return []HandType{断幺九}
	}
	return []HandType{}
}

func (hc HupaiCalculater) ipeko() []HandType {
	if hc.standard.IsMenzen() && hc.standard.SameMentsuVariationCountInMenzen() == 1 {
		return []HandType{一盃口}
	}
	return []HandType{}
}

func (hc HupaiCalculater) sansyokuDoujun() []HandType { panic("not implemented") }
func (hc HupaiCalculater) ittu() []HandType           { panic("not implemented") }
func (hc HupaiCalculater) chanta() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) toitoi() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) sanAnko() []HandType        { panic("not implemented") }
func (hc HupaiCalculater) sanKantsu() []HandType      { panic("not implemented") }
func (hc HupaiCalculater) sansyokuDoko() []HandType   { panic("not implemented") }

func (hc HupaiCalculater) honnro() []HandType {
	if !hc.standard.HasChunchan() && hc.standard.HasZi() {
		return []HandType{混老頭}
	}
	return []HandType{}
}

func (hc HupaiCalculater) syosangen() []HandType { panic("not implemented") }
func (hc HupaiCalculater) honnitsu() []HandType  { panic("not implemented") }
func (hc HupaiCalculater) junchan() []HandType   { panic("not implemented") }

func (hc HupaiCalculater) ryanpeko() []HandType {
	if hc.standard.IsMenzen() && hc.standard.SameMentsuVariationCountInMenzen() == 2 {
		return []HandType{二盃口}
	}
	return []HandType{}
}

func (hc HupaiCalculater) tinnitsu() []HandType  { panic("not implemented") }
func (hc HupaiCalculater) suAnko() []HandType    { panic("not implemented") }
func (hc HupaiCalculater) daisangen() []HandType { panic("not implemented") }
func (hc HupaiCalculater) sushi() []HandType     { panic("not implemented") }
func (hc HupaiCalculater) tsuiso() []HandType    { panic("not implemented") }
func (hc HupaiCalculater) ryuiso() []HandType    { panic("not implemented") }

func (hc HupaiCalculater) chinro() []HandType {
	if !hc.standard.HasChunchan() && !hc.standard.HasZi() {
		return []HandType{清老頭}
	}
	return []HandType{}
}

func (hc HupaiCalculater) suKantsu() []HandType {
	c := 0
	for _, m := range hc.standard.FiveBlocks() {
		if m.TypeIs(mentsuTypeKantsu) {
			c++
		}
	}

	if c == 4 {
		return []HandType{四槓子}
	}
	return []HandType{}
}

func (hc HupaiCalculater) churen() []HandType { panic("not implemented") }

type FullHupaiCalculater struct {
	fullParrern FullHoluPattern
	zhuangfeng  Zhuangfeng
	zifeng      Zifeng
	isTsumo     bool
}

func (fhc FullHupaiCalculater) Hupai() []AllHands {
	ret := []AllHands{}
	ret = append(ret, fhc.kokushi()) // 国士無双・国士無双十三面
	if len(ret) > 0 {
		return ret
	}

	ret = append(ret, fhc.titoitsuAll()) // 七対子(複合役も含め)

	for _, standard := range fhc.fullParrern.Standard {
		all := AllHands{}

		calculater := NewHupaiCalculater(standard, fhc.zhuangfeng, fhc.zifeng, fhc.isTsumo)

		all = append(all, calculater.tsumo()...)          // 門前清自摸和
		all = append(all, calculater.fengpai()...)        // 場風・自風・白・發・中
		all = append(all, calculater.pinfu()...)          // 平和
		all = append(all, calculater.tanyao()...)         // 断幺九
		all = append(all, calculater.ipeko()...)          // 一盃口
		all = append(all, calculater.sansyokuDoujun()...) // 三色同順
		all = append(all, calculater.ittu()...)           // 一気通貫
		all = append(all, calculater.chanta()...)         // 混全帯幺九
		all = append(all, calculater.toitoi()...)         // 対々和
		all = append(all, calculater.sanAnko()...)        // 三暗刻
		all = append(all, calculater.sanKantsu()...)      // 三槓子
		all = append(all, calculater.sansyokuDoko()...)   // 三色同刻
		all = append(all, calculater.honnro()...)         // 混老頭
		all = append(all, calculater.syosangen()...)      // 小三元
		all = append(all, calculater.honnitsu()...)       // 混一色
		all = append(all, calculater.junchan()...)        // 純全帯幺九
		all = append(all, calculater.ryanpeko()...)       // 二盃口
		all = append(all, calculater.tinnitsu()...)       // 清一色
		all = append(all, calculater.suAnko()...)         // 四暗刻・四暗刻単騎
		all = append(all, calculater.daisangen()...)      // 大三元
		all = append(all, calculater.sushi()...)          // 小四喜・大四喜
		all = append(all, calculater.tsuiso()...)         // 字一色
		all = append(all, calculater.ryuiso()...)         // 緑一色
		all = append(all, calculater.chinro()...)         // 清老頭
		all = append(all, calculater.suKantsu()...)       // 四槓子
		all = append(all, calculater.churen()...)         // 九蓮宝燈・純正九蓮宝燈
	}

	return ret
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

func (fhc FullHupaiCalculater) titoitsuAll() AllHands {
	titoitsu := fhc.fullParrern.Titoitsu

	if titoitsu == nil {
		return AllHands{}
	}

	ret := AllHands{七対子}

	// 門前清自摸和
	if fhc.isTsumo {
		ret = append(ret, 門前清自摸和)
	}

	/*
		断幺九
		混老頭
		混一色
		清一色
		字一色
	*/
	var (
		isTanyao     = true
		isHonroto    = true
		isTuisio     = true
		colorTypeMap = map[paiType]struct{}{}
		existZi      = false
	)
	for _, pai := range titoitsu.Menzen {
		if _, isYaojiu := YaojiuMap[pai]; isYaojiu {
			isTanyao = false
		} else {
			isHonroto = false
		}

		if pai.TypeIs(paiTypeZi) {
			existZi = true
		} else {
			isTuisio = false
			colorTypeMap[pai.Type] = struct{}{}
		}
	}
	if isTanyao {
		ret = append(ret, 断幺九)
	}
	if isHonroto {
		ret = append(ret, 混老頭)
	}
	if isTuisio {
		ret = append(ret, 字一色)
	}

	if len(colorTypeMap) == 1 {
		if existZi {
			ret = append(ret, 混一色)
		} else {
			ret = append(ret, 清一色)
		}
	}

	return ret
}
