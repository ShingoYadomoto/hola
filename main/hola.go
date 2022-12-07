package main

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

/*
function hule_mianzi_all(shoupai) {

    var mianzi = [[]];

    for (var s of ['m','p','s']) {
        var new_mianzi = [];
        var sub_mianzi = hule_mianzi(s, shoupai._shouli[s], 0);
        for (var m of mianzi) {
            for (var n of sub_mianzi) {
                new_mianzi.push(m.concat(n));
            }
        }
        mianzi = new_mianzi;
    }

    var sub_mianzi_z = [];
    for (var n = 1; n <= 7; n++) {
        if (shoupai._shouli.z[n-1] == 0) continue;
        if (shoupai._shouli.z[n-1] != 3) return [];
        sub_mianzi_z.push('z'+n+n+n);
    }

    for (var i = 0; i < mianzi.length; i++) {
        mianzi[i] = mianzi[i].concat(sub_mianzi_z)
                             .concat(shoupai._fulou);
    }

    return mianzi;
}

function add_hulepai(mianzi, hulepai) {

    var regexp   = new RegExp('^(' + hulepai[0] + '.*' + hulepai[1] +')');
    var replacer = '$1' + hulepai.substr(2) + '_';

    var add_mianzi = [];
    for (var i = 0; i < mianzi.length; i++) {
        if (mianzi[i].match(/[\-\+\=]/)) continue;
        if (i > 0 && mianzi[i] == mianzi[i-1]) continue;
        var rep = mianzi[i].replace(regexp, replacer);
        if (rep == mianzi[i]) continue;
        var new_mianzi = mianzi.concat();
        new_mianzi[i] = rep;
        add_mianzi.push(new_mianzi);
    }

    return add_mianzi;
}

function hule_yiban(shoupai, rongpai) {

    var hulepai = rongpai || shoupai._zimo;

    var hule_mianzi = [];
    for (var s in shoupai._shouli) {
        var pai = shoupai._shouli[s];
        for (var n = 1; n <= pai.length; n++) {
            if (pai[n-1] < 2) continue;
            var jiangpai = s+n+n;
            pai[n-1] -= 2;
            for (var mianzi of hule_mianzi_all(shoupai)) {
                mianzi.unshift(jiangpai);
                for (var add_mianzi of add_hulepai(mianzi, hulepai)) {
                    hule_mianzi.push(add_mianzi);
                }
            }
            pai[n-1] += 2;
        }
    }

    return hule_mianzi;
}

*/

// HolaYiban は4面子1雀頭形の処理。
// まず可能性のある雀頭を抜き出し(パターン3の処理)、hule_mianzi_all() に処理を任せた後、add_hulepai() を呼出して和了牌の位置を決めている。
// hule_mianzi() は各色の中で面子を探す処理(パターン2はここで処理)。
// hule_mianzi_all() は4面子1雀頭形について雀頭以外の面子を探す処理。各色ごとに hule_mianzi() を呼出している。
// add_hulepai() は和了牌を可能性のあるすべての面子に入れる処理(パターン4はここで処理)。
func HolaYiban(hand Hand, rongpai *pai) []StandardHolaHand { return nil }

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
