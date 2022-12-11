package hola_go

import (
	"sort"
)

// HoluMianzi は各色の中で面子を探す処理
func HoluMianzi(pais map[pai]int, checkPai pai) [][]mentsu {
	if checkPai.Index == 10 {
		return [][]mentsu{{}}
	}

	if pais[checkPai] == 0 {
		checkPai.Index++
		return HoluMianzi(pais, checkPai)
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
		shuntsu = HoluMianzi(pais, checkPai)
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
		kezi = HoluMianzi(pais, checkPai)
		pais[checkPai] += 3
		for i, _ := range kezi {
			kezi[i] = append(kezi[i], mentsu{pais: []pai{checkPai, checkPai, checkPai}})
		}
	}

	ret := append(shuntsu, kezi...)
	return ret
}

// HuleMianziAll は4面子1雀頭形について雀頭以外の面子を探す処理。各色ごとに hule_mianzi() を呼出している。
func HuleMianziAll(HoluPattern HoluPattern) [][]mentsu {
	mianzi := [][]mentsu{{}}

	for _, t := range []paiType{paiTypeManzu, paiTypePinzu, paiTypeSozu} {
		pais := map[pai]int{}
		for p, count := range HoluPattern.Menzen {
			if p.TypeIs(t) {
				pais[p] = count
			}
		}

		newMianzi := [][]mentsu{}
		unit := pai{Type: t, Index: 1}
		subMianzi := HoluMianzi(pais, unit)

		for _, m := range mianzi {
			for _, n := range subMianzi {
				newMianzi = append(newMianzi, append(m, n...))
			}
		}
		mianzi = newMianzi
	}

	subMianziZ := []mentsu{}
	for p, count := range HoluPattern.Menzen {
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

// HoluYiban は4面子1雀頭形の処理。
// まず可能性のある雀頭を抜き出し、hule_mianzi_all() に処理を任せた後、add_hulepai() を呼出して和了牌の位置を決めている。
func HoluYiban(holuPattern HoluPattern, rongpai *pai) []StandardHoluPattern {
	var HoluPai pai
	if rongpai == nil {
		HoluPai = *holuPattern.TsumoPai
	} else {
		HoluPai = *rongpai
	}

	var (
		huleMianzi = []StandardHoluPattern{}
		paiList    = make([]pai, len(holuPattern.Menzen))
	)
	i := 0
	for pai, _ := range holuPattern.Menzen {
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
		count := holuPattern.Menzen[pai]
		if count < 2 {
			continue
		}

		HoluHoluPattern := StandardHoluPattern{
			Head:        pai,
			HoluPai:     HoluPai,
			FulouMentsu: holuPattern.FulouMentsuList,
		}
		holuPattern.Menzen[pai] -= 2
		all := HuleMianziAll(holuPattern)
		for _, mianzi := range all {
			HoluHoluPattern.Mentsu = mianzi
			huleMianzi = append(huleMianzi, HoluHoluPattern)
		}
		holuPattern.Menzen[pai] += 2
	}

	return huleMianzi
}

// HoluGiduizi は七対子形の処理。
func HoluGiduizi(HoluPattern HoluPattern, rongpai *pai) *TitoitsuHoluPattern {
	if len(HoluPattern.FulouMentsuList) > 0 {
		return nil
	}

	var HoluPai pai
	if rongpai == nil {
		HoluPai = *HoluPattern.TsumoPai
	} else {
		HoluPai = *rongpai
	}

	ret := &TitoitsuHoluPattern{
		Menzen:  make([]pai, 7),
		HoluPai: HoluPai,
	}
	i := 0
	for pai, count := range HoluPattern.Menzen {
		if count != 2 {
			return nil
		}

		ret.Menzen[i] = pai
		i++
	}

	return ret
}

// HoluGuoshi は国士無双形の処理。
func HoluGuoshi(HoluPattern HoluPattern, rongpai *pai) *KokushiHoluPattern {
	if len(HoluPattern.FulouMentsuList) > 0 {
		return nil
	}

	var HoluPai pai
	if rongpai == nil {
		HoluPai = *HoluPattern.TsumoPai
	} else {
		HoluPai = *rongpai
	}

	ret := &KokushiHoluPattern{
		HoluPai: HoluPai,
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
	for pai, count := range HoluPattern.Menzen {
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

// Holu はメイン処理で、HoluYiban()、HoluGiduizi()、HoluGuoshi() を呼出してその結果を1つにまとめている
func Holu(HoluPattern HoluPattern, rongpai *pai) FullHoluPattern {
	if rongpai != nil {
		HoluPattern.Tsumo(*rongpai)
	}

	var (
		standard = HoluYiban(HoluPattern, rongpai)
		titoitsu = HoluGiduizi(HoluPattern, rongpai)
	)

	return FullHoluPattern{
		Standard: standard,
		Titoitsu: titoitsu,
		Kokushi:  HoluGuoshi(HoluPattern, rongpai),
		IsTsumo:  rongpai == nil,
	}
}

func Tsumo(HoluPattern HoluPattern) FullHoluPattern {
	return Holu(HoluPattern, nil)
}

func Rong(HoluPattern HoluPattern, rongpai *pai) FullHoluPattern {
	return Holu(HoluPattern, rongpai)
}
