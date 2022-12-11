package hola_go

type (
	paiType  int
	paiIndex int

	pai struct {
		Type  paiType
		Index paiIndex
	}
)

const (
	paiTypeManzu paiType = iota + 1
	paiTypePinzu
	paiTypeSozu
	paiTypeZi
)

const (
	paiIndex1 paiIndex = iota + 1
	paiIndex2
	paiIndex3
	paiIndex4
	paiIndex5
	paiIndex6
	paiIndex7
	paiIndex8
	paiIndex9
)

func (p pai) IsZero() bool {
	return p.Type == 0 || p.Index == 0
}

func (p pai) TypeIs(t paiType) bool {
	return p.Type == t
}

var (
	一萬 = pai{Type: paiTypeManzu, Index: paiIndex1}
	二萬 = pai{Type: paiTypeManzu, Index: paiIndex2}
	三萬 = pai{Type: paiTypeManzu, Index: paiIndex3}
	四萬 = pai{Type: paiTypeManzu, Index: paiIndex4}
	五萬 = pai{Type: paiTypeManzu, Index: paiIndex5}
	六萬 = pai{Type: paiTypeManzu, Index: paiIndex6}
	七萬 = pai{Type: paiTypeManzu, Index: paiIndex7}
	八萬 = pai{Type: paiTypeManzu, Index: paiIndex8}
	九萬 = pai{Type: paiTypeManzu, Index: paiIndex9}

	一筒 = pai{Type: paiTypePinzu, Index: paiIndex1}
	二筒 = pai{Type: paiTypePinzu, Index: paiIndex2}
	三筒 = pai{Type: paiTypePinzu, Index: paiIndex3}
	四筒 = pai{Type: paiTypePinzu, Index: paiIndex4}
	五筒 = pai{Type: paiTypePinzu, Index: paiIndex5}
	六筒 = pai{Type: paiTypePinzu, Index: paiIndex6}
	七筒 = pai{Type: paiTypePinzu, Index: paiIndex7}
	八筒 = pai{Type: paiTypePinzu, Index: paiIndex8}
	九筒 = pai{Type: paiTypePinzu, Index: paiIndex9}

	一索 = pai{Type: paiTypeSozu, Index: paiIndex1}
	二索 = pai{Type: paiTypeSozu, Index: paiIndex2}
	三索 = pai{Type: paiTypeSozu, Index: paiIndex3}
	四索 = pai{Type: paiTypeSozu, Index: paiIndex4}
	五索 = pai{Type: paiTypeSozu, Index: paiIndex5}
	六索 = pai{Type: paiTypeSozu, Index: paiIndex6}
	七索 = pai{Type: paiTypeSozu, Index: paiIndex7}
	八索 = pai{Type: paiTypeSozu, Index: paiIndex8}
	九索 = pai{Type: paiTypeSozu, Index: paiIndex9}

	東 = pai{Type: paiTypeZi, Index: paiIndex1}
	南 = pai{Type: paiTypeZi, Index: paiIndex2}
	西 = pai{Type: paiTypeZi, Index: paiIndex3}
	北 = pai{Type: paiTypeZi, Index: paiIndex4}
	白 = pai{Type: paiTypeZi, Index: paiIndex5}
	發 = pai{Type: paiTypeZi, Index: paiIndex6}
	中 = pai{Type: paiTypeZi, Index: paiIndex7}

	YaojiuMap = map[pai]struct{}{一萬: {}, 九萬: {}, 一筒: {}, 九筒: {}, 一索: {}, 九索: {}, 東: {}, 南: {}, 西: {}, 北: {}, 白: {}, 發: {}, 中: {}}
)
