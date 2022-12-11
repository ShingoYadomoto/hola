package hola_go

import (
	"math"
	"sort"
)

type (
	mentsu struct {
		pais    []pai
		holuPai *pai
	}
	mentsuType int
)

const (
	mentsuTypeShuntsu mentsuType = iota + 1
	mentsuTypeKotsu
	mentsuTypeKantsu
	mentsuTypehead
)

func newMentsu(pai1, pai2 pai, pais ...pai) mentsu {
	m := mentsu{pais: []pai{pai1, pai2}}
	for _, pai := range pais {
		m.pais = append(m.pais, pai)
	}
	return m
}

func NewShuntsu(pai1, pai2, pai3 pai) mentsu {
	return newMentsu(pai1, pai2, pai3)
}

func NewKotsu(pai pai) mentsu {
	return newMentsu(pai, pai, pai)
}

func NewKantsu(pai pai) mentsu {
	return newMentsu(pai, pai, pai, pai)
}

func NewHead(pai pai) mentsu {
	return newMentsu(pai, pai)
}

func (m mentsu) TypeIs(t mentsuType) bool {
	if len(m.pais) == 2 {
		return t == mentsuTypehead
	}
	if len(m.pais) == 4 {
		return t == mentsuTypeKantsu
	}

	prevent := pai{}
	for _, current := range m.pais {
		if !prevent.IsZero() && current != prevent {
			return t == mentsuTypeShuntsu
		}
		prevent = current
	}
	return t == mentsuTypeKotsu
}

func (m mentsu) Equal(compare mentsu) bool {
	if len(m.pais) != len(compare.pais) {
		return false
	}

	checker := map[pai]int{}
	for _, pai := range m.pais {
		checker[pai]++
	}

	for _, com := range compare.pais {
		if _, exist := checker[com]; !exist {
			return false
		}
		checker[com]--

		if checker[com] == 0 {
			delete(checker, com)
		}
	}

	return len(checker) == 0
}

func (m mentsu) HashCode() int {
	var (
		t    paiType
		code int
	)

	for _, pai := range m.pais {
		t = pai.Type
		code += int(math.Pow(3, float64(pai.Index)))
	}

	code += int(t) * 10000

	allTypes := map[mentsuType]int{
		mentsuTypehead:    1,
		mentsuTypeShuntsu: 2,
		mentsuTypeKotsu:   3,
		mentsuTypeKantsu:  4,
	}
	for t, c := range allTypes {
		if m.TypeIs(t) {
			code += c
			break
		}
	}
	return code
}

type (
	MentsuList []mentsu

	WaitType int
)

const (
	WaitTypeTanki WaitType = iota + 1
	WaitTypeTamen
	WaitTypeKanchan
	WaitTypeShanpon
)

func (ml MentsuList) WaitTypeIs(wt WaitType) bool {
	var (
		holuMentsu = mentsu{}
		holuPai    = pai{}
	)
	for _, m := range ml {
		if m.holuPai != nil {
			holuMentsu = m
			holuPai = *m.holuPai
			break
		}
	}

	if holuMentsu.TypeIs(mentsuTypehead) {
		return wt == WaitTypeTanki
	}
	if holuMentsu.TypeIs(mentsuTypeKotsu) {
		return wt == WaitTypeShanpon
	}

	paiList := holuMentsu.pais
	sort.Slice(paiList, func(i, j int) bool {
		return paiList[i].Index < paiList[j].Index
	})
	if paiList[1] == holuPai {
		return wt == WaitTypeKanchan
	}

	return wt == WaitTypeTamen
}

// 和了型全パターン
type FullHoluPattern struct {
	Standard []StandardHoluPattern
	Titoitsu *TitoitsuHoluPattern
	Kokushi  *KokushiHoluPattern
	IsTsumo  bool
}

// 4面子1雀頭系の和了型
type StandardHoluPattern struct {
	Mentsu      []mentsu        // 面子
	FulouMentsu FulouMentsuList // 副露面子
	Head        pai             // 雀頭
	HoluPai     pai             // 和了牌
}

func (shp StandardHoluPattern) FiveBlocks() []mentsu {
	ret := []mentsu{NewHead(shp.Head)}
	ret = append(ret, shp.Mentsu...)
	ret = append(ret, shp.FulouMentsu.MentsuList()...)

	return ret
}

func (shp StandardHoluPattern) IsMenzen() bool {
	return len(shp.FulouMentsu) == 0
}

func (shp StandardHoluPattern) HasSpecificKotsuOrKantsu(p pai) bool {
	for _, mentsu := range shp.FiveBlocks() {
		if mentsu.TypeIs(mentsuTypeKotsu) || mentsu.TypeIs(mentsuTypeKantsu) {
			if mentsu.pais[0] == p {
				return true
			}
		}
	}

	return false
}

func (shp StandardHoluPattern) HasYaojiu() bool {
	for _, mentsu := range shp.FiveBlocks() {
		for _, pai := range mentsu.pais {
			if _, isYaojiu := YaojiuMap[pai]; isYaojiu {
				return true
			}
		}
	}

	return false
}

func (shp StandardHoluPattern) HasChunchan() bool {
	for _, mentsu := range shp.FiveBlocks() {
		for _, pai := range mentsu.pais {
			if _, isYaojiu := YaojiuMap[pai]; !isYaojiu {
				return true
			}
		}
	}

	return false
}

func (shp StandardHoluPattern) HasZi() bool {
	for _, mentsu := range shp.FiveBlocks() {
		for _, pai := range mentsu.pais {
			if pai.TypeIs(paiTypeZi) {
				return true
			}
		}
	}

	return false
}

func (shp StandardHoluPattern) SameMentsuVariationCountInMenzen() int {
	uniqueMentsuCount := map[int]int{}
	for _, mentsu := range shp.Mentsu {
		hash := mentsu.HashCode()
		uniqueMentsuCount[hash]++
	}

	sameMentsuVariationCount := 0
	for _, count := range uniqueMentsuCount {
		if count > 1 {
			sameMentsuVariationCount++
		}
	}

	return sameMentsuVariationCount
}

func (shp StandardHoluPattern) IsZhuangfengpai(zhuangfeng Zhuangfeng) bool {
	var (
		fanpaimap = map[Zhuangfeng]pai{東場: 東, 西場: 西, 南場: 南, 北場: 北}
		fanpai    = fanpaimap[zhuangfeng]
	)

	return shp.HasSpecificKotsuOrKantsu(fanpai)
}

func (shp StandardHoluPattern) IsZifeng(zifeng Zifeng) bool {
	var (
		fanpaimap = map[Zifeng]pai{東家: 東, 西家: 西, 南家: 南, 北家: 北}
		fanpai    = fanpaimap[zifeng]
	)

	return shp.HasSpecificKotsuOrKantsu(fanpai)
}

// 七対子形の和了型
type TitoitsuHoluPattern struct {
	Menzen  []pai // 面前
	HoluPai pai   // 和了牌
}

// 国士無双形の和了型
type KokushiHoluPattern struct {
	Head    pai // 雀頭
	HoluPai pai // 和了牌
}

func (khp KokushiHoluPattern) IsDouble() bool {
	return khp.Head == khp.HoluPai
}

// 上がり時の手配構成
type HoluPattern struct {
	Menzen          map[pai]int
	TsumoPai        *pai
	FulouMentsuList []FulouMentsu
}

func (s HoluPattern) Tsumo(p pai) {
	s.TsumoPai = &p
	s.Menzen[p]++
}

type FulouMentsu struct {
	mentsu
	fulouPaiIndex int
	fulouFrom     Player
}

type FulouMentsuList []FulouMentsu

func (dml FulouMentsuList) MentsuList() []mentsu {
	l := make([]mentsu, len(dml))
	for i, fulouMentsu := range dml {
		l[i] = fulouMentsu.mentsu
	}
	return l
}
