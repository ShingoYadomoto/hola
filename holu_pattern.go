package hola_go

import (
	"math"
)

type (
	mentsu struct {
		pais []pai
	}
	mentsuType int
)

const (
	mentsuTypeShuntsu mentsuType = iota + 1
	mentsuTypeKotsu
	mentsuTypeKantsu
)

func (m mentsu) TypeIs(t mentsuType) bool {
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
		code += int(math.Pow(2, float64(pai.Index)))
	}

	code += int(t) * 10000
	return code
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

func (shp StandardHoluPattern) IsMenzen() bool {
	return len(shp.FulouMentsu) == 0
}

func (shp StandardHoluPattern) HasKotsu(p pai) bool {
	for _, mentsu := range append(shp.Mentsu, shp.FulouMentsu.MentsuList()...) {
		isKotsu := true
		for _, pai := range mentsu.pais {
			if pai != p {
				isKotsu = false
				break
			}
		}
		if isKotsu {
			return true
		}
	}

	return false
}

func (shp StandardHoluPattern) HasYaojiu() bool {
	for _, mentsu := range append(shp.Mentsu, shp.FulouMentsu.MentsuList()...) {
		for _, pai := range mentsu.pais {
			if _, isYaojiu := YaojiuMap[pai]; isYaojiu {
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

	return shp.HasKotsu(fanpai)
}

func (shp StandardHoluPattern) IsZifeng(zifeng Zifeng) bool {
	var (
		fanpaimap = map[Zifeng]pai{東家: 東, 西家: 西, 南家: 南, 北家: 北}
		fanpai    = fanpaimap[zifeng]
	)

	return shp.HasKotsu(fanpai)
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
