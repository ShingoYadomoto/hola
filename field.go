package hola_go

type Zhuangfeng int

const (
	東場 Zhuangfeng = iota + 1
	南場
	西場
	北場
)

func (z Zhuangfeng) Pai() pai {
	switch z {
	case 東場:
		return 東
	case 西場:
		return 西
	case 南場:
		return 南
	case 北場:
		return 北
	}
	return pai{}
}

type Zifeng int

const (
	東家 Zifeng = iota + 1
	南家
	西家
	北家
)

func (z Zifeng) Pai() pai {
	switch z {
	case 東家:
		return 東
	case 西家:
		return 西
	case 南家:
		return 南
	case 北家:
		return 北
	}
	return pai{}
}
