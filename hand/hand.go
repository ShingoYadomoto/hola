package hand

import (
	"github.com/ShingoYadomoto/hola-go/holu"
)

func Hupai(fullParrern holu.FullHoluPattern, zhuangfeng zhuangfeng, zifeng zifeng) AllHands {
	if fullParrern.Kokushi != nil {
		if fullParrern.Kokushi.IsDouble() {
			return []HandType{国士無双十三面}
		}
		return []HandType{国士無双}
	}

	return nil
}
