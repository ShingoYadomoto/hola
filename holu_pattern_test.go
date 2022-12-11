package hola_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mentsu_TypeIs(t *testing.T) {
	var (
		kantsu  = []pai{白, 白, 白, 白}
		kotsu   = []pai{一萬, 一萬, 一萬}
		shuntsu = []pai{一萬, 二萬, 三萬}
	)
	type fields struct {
		pais []pai
	}
	tests := []struct {
		name   string
		fields fields
		arg    mentsuType
		want   bool
	}{
		{name: "槓子 true", fields: fields{pais: kantsu}, arg: mentsuTypeKantsu, want: true},
		{name: "槓子 false1", fields: fields{pais: kotsu}, arg: mentsuTypeKantsu, want: false},
		{name: "槓子 false2", fields: fields{pais: shuntsu}, arg: mentsuTypeKantsu, want: false},

		{name: "刻子 true", fields: fields{pais: kotsu}, arg: mentsuTypeKotsu, want: true},
		{name: "刻子 false1", fields: fields{pais: kantsu}, arg: mentsuTypeKotsu, want: false},
		{name: "刻子 false2", fields: fields{pais: shuntsu}, arg: mentsuTypeKotsu, want: false},

		{name: "順子 true", fields: fields{pais: shuntsu}, arg: mentsuTypeShuntsu, want: true},
		{name: "順子 false1", fields: fields{pais: kantsu}, arg: mentsuTypeShuntsu, want: false},
		{name: "順子 false2", fields: fields{pais: kotsu}, arg: mentsuTypeShuntsu, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mentsu{
				pais: tt.fields.pais,
			}
			assert.Equalf(t, tt.want, m.TypeIs(tt.arg), "Type()")
		})
	}
}

func Test_mentsu_Equal(t *testing.T) {
	tests := []struct {
		name   string
		fields mentsu
		args   mentsu
		want   bool
	}{
		{
			name:   "true case. shuntsu",
			fields: mentsu{pais: []pai{一萬, 二萬, 三萬}},
			args:   mentsu{pais: []pai{一萬, 二萬, 三萬}},
			want:   true,
		},
		{
			name:   "true case. kotsu",
			fields: mentsu{pais: []pai{一萬, 一萬, 一萬}},
			args:   mentsu{pais: []pai{一萬, 一萬, 一萬}},
			want:   true,
		},
		{
			name:   "true case. kantsu",
			fields: mentsu{pais: []pai{一萬, 一萬, 一萬, 一萬}},
			args:   mentsu{pais: []pai{一萬, 一萬, 一萬, 一萬}},
			want:   true,
		},
		{
			name:   "true case. not same order",
			fields: mentsu{pais: []pai{三萬, 一萬, 二萬}},
			args:   mentsu{pais: []pai{一萬, 二萬, 三萬}},
			want:   true,
		},
		{
			name:   "false case",
			fields: mentsu{pais: []pai{一萬, 三萬, 二萬}},
			args:   mentsu{pais: []pai{一索, 二索, 三索}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mentsu{
				pais: tt.fields.pais,
			}
			assert.Equalf(t, tt.want, m.Equal(tt.args), "Equal(%v)", tt.args)
		})
	}
}

func Test_mentsu_HashCode(t *testing.T) {
	allMentsu := [][]pai{
		// 萬子
		{一萬, 二萬, 三萬},
		{二萬, 三萬, 四萬},
		{三萬, 四萬, 五萬},
		{四萬, 五萬, 六萬},
		{五萬, 六萬, 七萬},
		{六萬, 七萬, 八萬},
		{七萬, 八萬, 九萬},

		{一萬, 一萬, 一萬},
		{二萬, 二萬, 二萬},
		{三萬, 三萬, 三萬},
		{四萬, 四萬, 四萬},
		{五萬, 五萬, 五萬},
		{六萬, 六萬, 六萬},
		{七萬, 七萬, 七萬},
		{八萬, 八萬, 八萬},
		{九萬, 九萬, 九萬},

		{一萬, 一萬, 一萬, 一萬},
		{二萬, 二萬, 二萬, 二萬},
		{三萬, 三萬, 三萬, 三萬},
		{四萬, 四萬, 四萬, 四萬},
		{五萬, 五萬, 五萬, 五萬},
		{六萬, 六萬, 六萬, 六萬},
		{七萬, 七萬, 七萬, 七萬},
		{八萬, 八萬, 八萬, 八萬},
		{九萬, 九萬, 九萬, 九萬},

		{一萬, 一萬},
		{二萬, 二萬},
		{三萬, 三萬},
		{四萬, 四萬},
		{五萬, 五萬},
		{六萬, 六萬},
		{七萬, 七萬},
		{八萬, 八萬},
		{九萬, 九萬},

		// 萬子
		{一筒, 二筒, 三筒},
		{二筒, 三筒, 四筒},
		{三筒, 四筒, 五筒},
		{四筒, 五筒, 六筒},
		{五筒, 六筒, 七筒},
		{六筒, 七筒, 八筒},
		{七筒, 八筒, 九筒},

		{一筒, 一筒, 一筒},
		{二筒, 二筒, 二筒},
		{三筒, 三筒, 三筒},
		{四筒, 四筒, 四筒},
		{五筒, 五筒, 五筒},
		{六筒, 六筒, 六筒},
		{七筒, 七筒, 七筒},
		{八筒, 八筒, 八筒},
		{九筒, 九筒, 九筒},

		{一筒, 一筒, 一筒, 一筒},
		{二筒, 二筒, 二筒, 二筒},
		{三筒, 三筒, 三筒, 三筒},
		{四筒, 四筒, 四筒, 四筒},
		{五筒, 五筒, 五筒, 五筒},
		{六筒, 六筒, 六筒, 六筒},
		{七筒, 七筒, 七筒, 七筒},
		{八筒, 八筒, 八筒, 八筒},
		{九筒, 九筒, 九筒, 九筒},

		{一筒, 一筒},
		{二筒, 二筒},
		{三筒, 三筒},
		{四筒, 四筒},
		{五筒, 五筒},
		{六筒, 六筒},
		{七筒, 七筒},
		{八筒, 八筒},
		{九筒, 九筒},

		// 索子
		{一索, 二索, 三索},
		{二索, 三索, 四索},
		{三索, 四索, 五索},
		{四索, 五索, 六索},
		{五索, 六索, 七索},
		{六索, 七索, 八索},
		{七索, 八索, 九索},

		{一索, 一索, 一索},
		{二索, 二索, 二索},
		{三索, 三索, 三索},
		{四索, 四索, 四索},
		{五索, 五索, 五索},
		{六索, 六索, 六索},
		{七索, 七索, 七索},
		{八索, 八索, 八索},
		{九索, 九索, 九索},

		{一索, 一索, 一索, 一索},
		{二索, 二索, 二索, 二索},
		{三索, 三索, 三索, 三索},
		{四索, 四索, 四索, 四索},
		{五索, 五索, 五索, 五索},
		{六索, 六索, 六索, 六索},
		{七索, 七索, 七索, 七索},
		{八索, 八索, 八索, 八索},
		{九索, 九索, 九索, 九索},

		{一索, 一索},
		{二索, 二索},
		{三索, 三索},
		{四索, 四索},
		{五索, 五索},
		{六索, 六索},
		{七索, 七索},
		{八索, 八索},
		{九索, 九索},

		// 字牌
		{東, 東, 東},
		{南, 南, 南},
		{西, 西, 西},
		{北, 北, 北},
		{白, 白, 白},
		{發, 發, 發},
		{中, 中, 中},

		{東, 東, 東, 東},
		{南, 南, 南, 南},
		{西, 西, 西, 西},
		{北, 北, 北, 北},
		{白, 白, 白, 白},
		{發, 發, 發, 發},
		{中, 中, 中, 中},

		{東, 東},
		{南, 南},
		{西, 西},
		{北, 北},
		{白, 白},
		{發, 發},
		{中, 中},
	}

	checker := map[int][]pai{}
	for _, pais := range allMentsu {
		t.Run("", func(t *testing.T) {
			var (
				m    = mentsu{pais: pais}
				code = m.HashCode()
			)
			if pais1, dup := checker[code]; dup {
				t.Errorf("HashCode() = %v, mentsu1 %v,mentsu2 %v", code, pais1, pais)
			}
			checker[code] = pais
		})
	}
}
