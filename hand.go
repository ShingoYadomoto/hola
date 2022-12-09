package hola_go

func Hupai(fullParrern FullHoluPattern, zhuangfeng Zhuangfeng, zifeng Zifeng) AllHands {
	if fullParrern.Kokushi != nil {
		if fullParrern.Kokushi.IsDouble() {
			return []HandType{国士無双十三面}
		}
		return []HandType{国士無双}
	}

	menzen := true
	for _, s := range fullParrern.Standard {
		menzen = s.IsMenzen()
		break
	}
	_ = menzen

	return nil
}
