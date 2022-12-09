package holu

type (
	mentsu struct {
		pais []pai
	}

	// 和了型全パターン
	FullHoluPattern struct {
		Standard []StandardHoluPattern
		Titoitsu *TitoitsuHoluPattern
		Kokushi  *KokushiHoluPattern
		isTsumo  bool
	}

	// 4面子1雀頭系の和了型
	StandardHoluPattern struct {
		Mentsu      []mentsu // 面子
		FulouMentsu []mentsu // 副露面子
		Head        pai      // 雀頭
		HoluPai     pai      // 和了牌
	}

	// 七対子形の和了型
	TitoitsuHoluPattern struct {
		Menzen  []pai // 面前
		HoluPai pai   // 和了牌
	}

	// 国士無双形の和了型
	KokushiHoluPattern struct {
		Head    pai // 雀頭
		HoluPai pai // 和了牌
	}
)

func (shp StandardHoluPattern) IsMenzen() bool {
	return len(shp.FulouMentsu) == 0
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
