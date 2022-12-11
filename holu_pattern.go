package hola_go

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

func (m mentsu) Type() mentsuType {
	if len(m.pais) == 4 {
		return mentsuTypeKantsu
	}

	prevent := pai{}
	for _, current := range m.pais {
		if !prevent.IsZero() && current != prevent {
			return mentsuTypeShuntsu
		}
		prevent = current
	}
	return mentsuTypeKotsu
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

// 和了型全パターン
type FullHoluPattern struct {
	Standard []StandardHoluPattern
	Titoitsu *TitoitsuHoluPattern
	Kokushi  *KokushiHoluPattern
	IsTsumo  bool
}

// 4面子1雀頭系の和了型
type StandardHoluPattern struct {
	Mentsu      []mentsu // 面子
	FulouMentsu []mentsu // 副露面子
	Head        pai      // 雀頭
	HoluPai     pai      // 和了牌
}

func (shp StandardHoluPattern) IsMenzen() bool {
	return len(shp.FulouMentsu) == 0
}

func (shp StandardHoluPattern) HasKotsu(p pai) bool {
	for _, mentsu := range append(shp.Mentsu, shp.FulouMentsu...) {
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
	for _, mentsu := range append(shp.Mentsu, shp.FulouMentsu...) {
		for _, pai := range mentsu.pais {
			if _, isYaojiu := YaojiuMap[pai]; isYaojiu {
				return true
			}
		}
	}

	return false
}

//func (shp StandardHoluPattern) SameShuntsuCountInMenzen() int {
//	for i, mentsu1 := range shp.Mentsu {
//		for j, mentsu2 := range shp.Mentsu {
//			if i == j {
//				continue
//			}
//			if reflect.DeepEqual(mentsu1, mentsu2) {
//				return true
//			}
//		}
//	}
//
//	return false
//}

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
	Menzen   map[pai]int
	TsumoPai *pai
	FulouPai []Fulou
}

func (s HoluPattern) Tsumo(p pai) {
	s.TsumoPai = &p
	s.Menzen[p]++
}

type Fulou struct {
	pais          []pai
	fulouPaiIndex int
	fulouFrom     Player
}
