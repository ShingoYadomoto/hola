package hola_go

import (
	"fmt"
	"strings"
)

type HandType int

// AllHands 役の一覧
// ex:
// AllHands = []HandType{門前清自摸和, 翻牌場風}
type AllHands []HandType

// PossibleAllHands 手配の役全パターン
// ex:
// PossibleAllHands = []AllHands{
//   []HandType{七対子, 断幺九},
//   []HandType{二盃口, 断幺九},
// }
type PossibleAllHands []AllHands

func (pah PossibleAllHands) String() string {
	ret := ""
	for _, v1 := range pah {
		s := make([]string, len(v1))
		for i, v2 := range v1 {
			s[i] = v2.String()
		}
		ret += fmt.Sprintf("%s\n", strings.Join(s, ", "))
	}
	return ret
}

const (
	門前清自摸和 HandType = iota + 1
	翻牌場風
	翻牌自風
	翻牌白
	翻牌發
	翻牌中
	平和
	断幺九
	一盃口
	三色同順
	一気通貫
	混全帯幺九
	七対子
	対々和
	三暗刻
	三槓子
	三色同刻
	混老頭
	小三元
	混一色
	純全帯幺九
	二盃口
	清一色
	国士無双
	国士無双十三面
	四暗刻
	四暗刻単騎
	大三元
	小四喜
	大四喜
	字一色
	緑一色
	清老頭
	四槓子
	九蓮宝燈
	純正九蓮宝燈
)

type Fanshu int

func (ht HandType) Fanshu(isMenzen bool) Fanshu {
	var fanshu Fanshu

	switch ht {
	case 門前清自摸和, 翻牌場風, 翻牌自風, 翻牌白, 翻牌發, 翻牌中, 平和, 断幺九, 一盃口:
		fanshu = 1
	case 三色同順, 一気通貫, 混全帯幺九, 七対子, 対々和, 三暗刻, 三槓子, 三色同刻, 混老頭, 小三元, 混一色:
		fanshu = 2
	case 純全帯幺九, 二盃口:
		fanshu = 3
	case 清一色:
		fanshu = 6
	case 国士無双, 四暗刻, 大三元, 小四喜, 字一色, 緑一色, 清老頭, 四槓子, 九蓮宝燈:
		fanshu = 13
	case 大四喜, 国士無双十三面, 四暗刻単騎, 純正九蓮宝燈:
		fanshu = 26
	}

	if isMenzen {
		return fanshu
	}

	switch ht {
	case 三色同順, 一気通貫, 混全帯幺九:
		fanshu = 1
	case 混一色, 純全帯幺九:
		fanshu = 2
	case 清一色:
		fanshu = 5
	}

	return fanshu
}

func (ht HandType) String() string {
	m := map[HandType]string{
		門前清自摸和: "門前清自摸和", 翻牌場風: "翻牌場風", 翻牌自風: "翻牌自風", 翻牌白: "翻牌白", 翻牌發: "翻牌發", 翻牌中: "翻牌中", 平和: "平和", 断幺九: "断幺九", 一盃口: "一盃口",
		三色同順: "三色同順", 一気通貫: "一気通貫", 混全帯幺九: "混全帯幺九", 七対子: "七対子", 対々和: "対々和", 三暗刻: "三暗刻", 三槓子: "三槓子", 三色同刻: "三色同刻", 混老頭: "混老頭",
		小三元: "小三元", 混一色: "混一色", 純全帯幺九: "純全帯幺九", 二盃口: "二盃口", 清一色: "清一色", 国士無双: "国士無双", 国士無双十三面: "国士無双十三面", 四暗刻: "四暗刻",
		四暗刻単騎: "四暗刻単騎", 大三元: "大三元", 小四喜: "小四喜", 大四喜: "大四喜", 字一色: "字一色", 緑一色: "緑一色", 清老頭: "清老頭", 四槓子: "四槓子", 九蓮宝燈: "九蓮宝燈", 純正九蓮宝燈: "純正九蓮宝燈",
	}

	return m[ht]
}
