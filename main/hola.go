package main

import (
	"sort"
)

// HolaMianzi() は各色の中で面子を探す処理(パターン2はここで処理)。
func HolaMianzi(pais map[pai]int, checkPai pai) [][]mentsu {
	if checkPai.Index == 10 {
		return [][]mentsu{{}}
	}

	if pais[checkPai] == 0 {
		checkPai.Index++
		return HolaMianzi(pais, checkPai)
	}

	var (
		shuntsu     = [][]mentsu{}
		nextPai     = pai{Type: checkPai.Type, Index: checkPai.Index + 1}
		nextNextPai = pai{Type: checkPai.Type, Index: checkPai.Index + 2}
	)
	if checkPai.Index < paiIndex8 && pais[checkPai] > 0 && pais[nextPai] > 0 && pais[nextNextPai] > 0 {
		pais[checkPai]--
		pais[nextPai]--
		pais[nextNextPai]--
		shuntsu = HolaMianzi(pais, checkPai)
		pais[checkPai]++
		pais[nextPai]++
		pais[nextNextPai]++
		for i, _ := range shuntsu {
			shuntsu[i] = append(shuntsu[i], mentsu{pais: []pai{checkPai, nextPai, nextNextPai}})
		}
	}

	var (
		kezi = [][]mentsu{}
	)
	if pais[checkPai] >= 3 {
		pais[checkPai] -= 3
		kezi = HolaMianzi(pais, checkPai)
		pais[checkPai] += 3
		for i, _ := range kezi {
			kezi[i] = append(kezi[i], mentsu{pais: []pai{checkPai, checkPai, checkPai}})
		}
	}

	ret := append(shuntsu, kezi...)
	return ret
}

// HuleMianziAll() は4面子1雀頭形について雀頭以外の面子を探す処理。各色ごとに hule_mianzi() を呼出している。
func HuleMianziAll(hand Hand) [][]mentsu {
	mianzi := [][]mentsu{{}}

	for _, t := range []paiType{paiTypeManzu, paiTypePinzu, paiTypeSozu} {
		pais := map[pai]int{}
		for p, count := range hand.Menzen {
			if p.TypeIs(t) {
				pais[p] = count
			}
		}

		newMianzi := [][]mentsu{}
		unit := pai{Type: t, Index: 1}
		subMianzi := HolaMianzi(pais, unit)

		for _, m := range mianzi {
			for _, n := range subMianzi {
				newMianzi = append(newMianzi, append(m, n...))
			}
		}
		mianzi = newMianzi
	}

	subMianziZ := []mentsu{}
	for p, count := range hand.Menzen {
		if p.TypeIs(paiTypeZi) {
			switch count {
			case 0:
				continue
			case 3:
				subMianziZ = append(subMianziZ, mentsu{pais: []pai{p, p, p}})
			default:
				return [][]mentsu{}
			}
		}
	}

	for i, m := range mianzi {
		mianzi[i] = append(m, subMianziZ...)
	}

	return mianzi
}

// HolaYiban は4面子1雀頭形の処理。
// まず可能性のある雀頭を抜き出し(パターン3の処理)、hule_mianzi_all() に処理を任せた後、add_hulepai() を呼出して和了牌の位置を決めている。
func HolaYiban(hand Hand, rongpai *pai) []StandardHolaHand {
	var holaPai pai
	if rongpai == nil {
		holaPai = *hand.TsumoPai
	} else {
		holaPai = *rongpai
	}

	var (
		huleMianzi = []StandardHolaHand{}
		paiList    = make([]pai, len(hand.Menzen))
	)
	i := 0
	for pai, _ := range hand.Menzen {
		paiList[i] = pai
		i++
	}

	sort.Slice(paiList, func(i, j int) bool {
		if paiList[i].Type == paiList[j].Type {
			return paiList[i].Index < paiList[j].Index
		}

		return paiList[i].Type < paiList[j].Type
	})
	for _, pai := range paiList {
		count := hand.Menzen[pai]
		if count < 2 {
			continue
		}

		holahand := StandardHolaHand{
			Head:    pai,
			HolaPai: holaPai,
		}
		hand.Menzen[pai] -= 2
		all := HuleMianziAll(hand)
		for _, mianzi := range all {
			holahand.Mentsu = mianzi
			huleMianzi = append(huleMianzi, holahand)
		}
		hand.Menzen[pai] += 2
	}

	return huleMianzi
}

// HolaGiduizi は七対子形の処理。
func HolaGiduizi(hand Hand, rongpai *pai) *TitoitsuHolaHand {
	if len(hand.FulouPai) > 0 {
		return nil
	}

	var holaPai pai
	if rongpai == nil {
		holaPai = *hand.TsumoPai
	} else {
		holaPai = *rongpai
	}

	ret := &TitoitsuHolaHand{
		Hand:    make([]pai, 7),
		HolaPai: holaPai,
	}
	i := 0
	for pai, count := range hand.Menzen {
		if count != 2 {
			return nil
		}

		ret.Hand[i] = pai
		i++
	}

	return ret
}

// HolaGuoshi は国士無双形の処理。
func HolaGuoshi(hand Hand, rongpai *pai) *KokushiHolaHand {
	if len(hand.FulouPai) > 0 {
		return nil
	}

	var holaPai pai
	if rongpai == nil {
		holaPai = *hand.TsumoPai
	} else {
		holaPai = *rongpai
	}

	ret := &KokushiHolaHand{
		HolaPai: holaPai,
	}

	required := map[pai]struct{}{
		一萬: {},
		九萬: {},
		一筒: {},
		九筒: {},
		一索: {},
		九索: {},
		東:  {},
		南:  {},
		西:  {},
		北:  {},
		白:  {},
		發:  {},
		中:  {},
	}
	for pai, count := range hand.Menzen {
		if _, ok := required[pai]; !ok {
			return nil
		}

		switch pai.Type {
		case paiTypeManzu, paiTypePinzu, paiTypeSozu:
			if pai.Index != paiIndex1 && pai.Index != paiIndex9 {
				return nil
			}
		}

		if count == 2 {
			ret.Head = pai
		}

		delete(required, pai)
	}

	if len(required) != 0 {
		return nil
	}

	return ret
}

/*
// hola() はメイン処理で、HolaYiban()、HolaGiduizi()、HolaGuoshi() を呼出してその結果を1つにまとめている(パターン1の処理)。
func hola(hand Hand, rongpai *pai) []HolaHand {
	if rongpai != nil {
		hand.Tsumo(*rongpai)
	}

	return append(HolaYiban(hand, rongpai), HolaGiduizi(hand, rongpai), HolaGuoshi(hand, rongpai))
}

func Tsumo(hand Hand) []HolaHand {
	return hola(hand, nil)
}

func Rong(hand Hand, rongpai *pai) []HolaHand {
	return hola(hand, rongpai)
}
*/
