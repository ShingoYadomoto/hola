package main

type (
	mentsu struct {
		pais []pai
	}

	// 4面子1雀頭系の和了型
	StandardHolaHand struct {
		Mentsu  []mentsu // 面子
		Head    pai      // 雀頭
		HolaPai pai      // 和了牌
	}

	// 七対子形の和了型
	TitoitsuHolaHand struct {
		Hand    []pai // 面前
		HolaPai pai   // 和了牌
	}

	// 国士無双形の和了型
	KokushiHolaHand struct {
		Head    pai // 雀頭
		HolaPai pai // 和了牌
	}
)

// 上がり時の手配構成
type Hand struct {
	Menzen   map[pai]int
	TsumoPai *pai
	FulouPai []Fulou
}

func (s Hand) Tsumo(p pai) {
	s.TsumoPai = &p
	s.Menzen[p]++
}

type Fulou struct {
	paiList       [4]pai
	fulouPaiIndex int
	fulouFrom     Player
}
