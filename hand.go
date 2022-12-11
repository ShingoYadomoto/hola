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

func (hc HupaiCalculater) menzen() []HandType {
	if hc.standard.IsMenzen() {
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

func (hc HupaiCalculater) pinfu() []HandType { panic("not implemented") }

func (hc HupaiCalculater) tanyao() []HandType {
	if !hc.standard.HasYaojiu() {
		return []HandType{断幺九}
	}
	return []HandType{}
}

func (hc HupaiCalculater) ipeko() []HandType          { panic("not implemented") }
func (hc HupaiCalculater) sansyokuDoujun() []HandType { panic("not implemented") }
func (hc HupaiCalculater) ittu() []HandType           { panic("not implemented") }
func (hc HupaiCalculater) chanta() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) toitoi() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) sanAnko() []HandType        { panic("not implemented") }
func (hc HupaiCalculater) sanKantsu() []HandType      { panic("not implemented") }
func (hc HupaiCalculater) sansyokuDoko() []HandType   { panic("not implemented") }
func (hc HupaiCalculater) honro() []HandType          { panic("not implemented") }
func (hc HupaiCalculater) syosangen() []HandType      { panic("not implemented") }
func (hc HupaiCalculater) honitsu() []HandType        { panic("not implemented") }
func (hc HupaiCalculater) junchan() []HandType        { panic("not implemented") }
func (hc HupaiCalculater) ryanpeko() []HandType       { panic("not implemented") }
func (hc HupaiCalculater) tinitsu() []HandType        { panic("not implemented") }
func (hc HupaiCalculater) suAnko() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) daisangen() []HandType      { panic("not implemented") }
func (hc HupaiCalculater) sushi() []HandType          { panic("not implemented") }
func (hc HupaiCalculater) tsuiso() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) ryuiso() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) chinro() []HandType         { panic("not implemented") }
func (hc HupaiCalculater) suKantsu() []HandType       { panic("not implemented") }
func (hc HupaiCalculater) churen() []HandType         { panic("not implemented") }

type FullHupaiCalculater struct {
	fullParrern FullHoluPattern
	zhuangfeng  Zhuangfeng
	zifeng      Zifeng
}

func (fhc FullHupaiCalculater) Hupai() AllHands {
	all := AllHands{}
	all = append(all, fhc.kokushi()...)  //	国士無双・国士無双十三面
	all = append(all, fhc.titoitsu()...) // 七対子
	for _, standard := range fhc.fullParrern.Standard {
		calculater := NewHupaiCalculater(standard, fhc.zhuangfeng, fhc.zifeng)

		all = append(all, calculater.menzen()...)         // 門前
		all = append(all, calculater.fengpai()...)        // 場風・自風・白・發・中
		all = append(all, calculater.pinfu()...)          //平和
		all = append(all, calculater.tanyao()...)         //断幺九
		all = append(all, calculater.ipeko()...)          //一盃口
		all = append(all, calculater.sansyokuDoujun()...) //三色同順
		all = append(all, calculater.ittu()...)           //一気通貫
		all = append(all, calculater.chanta()...)         //混全帯幺九
		all = append(all, calculater.toitoi()...)         //対々和
		all = append(all, calculater.sanAnko()...)        //三暗刻
		all = append(all, calculater.sanKantsu()...)      //三槓子
		all = append(all, calculater.sansyokuDoko()...)   //三色同刻
		all = append(all, calculater.honro()...)          //混老頭
		all = append(all, calculater.syosangen()...)      //小三元
		all = append(all, calculater.honitsu()...)        //混一色
		all = append(all, calculater.junchan()...)        //純全帯幺九
		all = append(all, calculater.ryanpeko()...)       //二盃口
		all = append(all, calculater.tinitsu()...)        //清一色
		all = append(all, calculater.suAnko()...)         //四暗刻・四暗刻単騎
		all = append(all, calculater.daisangen()...)      //大三元
		all = append(all, calculater.sushi()...)          //小四喜・大四喜
		all = append(all, calculater.tsuiso()...)         //字一色
		all = append(all, calculater.ryuiso()...)         //緑一色
		all = append(all, calculater.chinro()...)         //清老頭
		all = append(all, calculater.suKantsu()...)       //四槓子
		all = append(all, calculater.churen()...)         //九蓮宝燈・純正九蓮宝燈
	}

	return all
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
