package holu

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHoluGuoshi(t *testing.T) {
	type args struct {
		holuPattern HoluPattern
		rongpai     *pai
	}
	tests := []struct {
		name string
		args args
		want *KokushiHoluPattern
	}{
		{
			name: "tsumo",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 1,
						九萬: 1,
						一筒: 1,
						九筒: 1,
						一索: 1,
						九索: 1,
						東:  1,
						南:  1,
						西:  1,
						北:  1,
						白:  1,
						發:  1,
						中:  2,
					},
					TsumoPai: &發,
				},
			},
			want: &KokushiHoluPattern{
				Head:    中,
				HoluPai: 發,
			},
		},

		{
			name: "tsumo. 13面",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 1,
						九萬: 1,
						一筒: 1,
						九筒: 1,
						一索: 1,
						九索: 1,
						東:  1,
						南:  1,
						西:  1,
						北:  1,
						白:  1,
						發:  1,
						中:  2,
					},
					TsumoPai: &中,
				},
			},
			want: &KokushiHoluPattern{
				Head:    中,
				HoluPai: 中,
			},
		},

		{
			name: "rong",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 1,
						九萬: 1,
						一筒: 1,
						九筒: 1,
						一索: 1,
						九索: 1,
						東:  1,
						南:  1,
						西:  1,
						北:  1,
						白:  1,
						發:  1,
						中:  2,
					},
				},
				rongpai: &發,
			},
			want: &KokushiHoluPattern{
				Head:    中,
				HoluPai: 發,
			},
		},

		{
			name: "rong. 13面",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 1,
						九萬: 1,
						一筒: 1,
						九筒: 1,
						一索: 1,
						九索: 1,
						東:  1,
						南:  1,
						西:  1,
						北:  1,
						白:  1,
						發:  1,
						中:  2,
					},
				},
				rongpai: &中,
			},
			want: &KokushiHoluPattern{
				Head:    中,
				HoluPai: 中,
			},
		},

		{
			name: "not kokushi",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 3,
						九萬: 3,
						一筒: 3,
						九筒: 3,
						一索: 2,
					},
				},
				rongpai: &一索,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HoluGuoshi(tt.args.holuPattern, tt.args.rongpai); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HoluGuoshi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoluGiduizi(t *testing.T) {
	type args struct {
		holuPattern HoluPattern
		rongpai     *pai
	}
	tests := []struct {
		name string
		args args
		want *TitoitsuHoluPattern
	}{
		{
			name: "tsumo",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 2,
						九萬: 2,
						一筒: 2,
						九筒: 2,
						一索: 2,
						九索: 2,
						東:  2,
					},
					TsumoPai: &發,
				},
			},
			want: &TitoitsuHoluPattern{
				Menzen:  []pai{一萬, 九萬, 一筒, 九筒, 一索, 九索, 東},
				HoluPai: 發,
			},
		},

		{
			name: "rong",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 2,
						九萬: 2,
						一筒: 2,
						九筒: 2,
						一索: 2,
						九索: 2,
						東:  2,
					},
				},
				rongpai: &發,
			},
			want: &TitoitsuHoluPattern{
				Menzen:  []pai{一萬, 九萬, 一筒, 九筒, 一索, 九索, 東},
				HoluPai: 發,
			},
		},

		{
			name: "not Titoitsu",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 1,
						九萬: 1,
						一筒: 1,
						九筒: 1,
						一索: 1,
						九索: 1,
						東:  1,
						南:  1,
						西:  1,
						北:  1,
						白:  1,
						發:  1,
						中:  2,
					},
				},
				rongpai: &中,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HoluGiduizi(tt.args.holuPattern, tt.args.rongpai)
			if tt.want != nil {
				assert.ElementsMatch(t, got.Menzen, tt.want.Menzen)
				assert.Equal(t, got.HoluPai, tt.want.HoluPai)
			} else {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestHoluMianzi(t *testing.T) {
	type args struct {
		pais     map[pai]int
		checkPai pai
	}
	tests := []struct {
		name string
		args args
		want [][]mentsu
	}{
		{
			args: args{
				pais:     map[pai]int{一萬: 1, 二萬: 3, 三萬: 3},
				checkPai: 一萬,
			},
			want: [][]mentsu{},
		},
		{
			args: args{
				pais:     map[pai]int{七筒: 1, 八筒: 1, 九筒: 3},
				checkPai: 一筒,
			},
			want: [][]mentsu{},
		},
		{
			args: args{
				pais:     map[pai]int{},
				checkPai: 一索,
			},
			want: [][]mentsu{},
		},
		{
			args: args{
				pais:     map[pai]int{一萬: 3, 二萬: 1, 三萬: 3},
				checkPai: 一萬,
			},
			want: [][]mentsu{},
		},
		{
			args: args{
				pais:     map[pai]int{一萬: 3, 二萬: 3, 三萬: 1},
				checkPai: 一萬,
			},
			want: [][]mentsu{},
		},
		{
			args: args{
				pais:     map[pai]int{一萬: 3, 二萬: 3, 三萬: 3},
				checkPai: 一萬,
			},
			want: [][]mentsu{
				{
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{一萬, 二萬, 三萬}},
				},
				{
					{pais: []pai{三萬, 三萬, 三萬}},
					{pais: []pai{二萬, 二萬, 二萬}},
					{pais: []pai{一萬, 一萬, 一萬}},
				},
			},
		},
		{
			args: args{
				pais:     map[pai]int{七筒: 1, 八筒: 1, 九筒: 1},
				checkPai: 一筒,
			},
			want: [][]mentsu{
				{
					{pais: []pai{七筒, 八筒, 九筒}},
				},
			},
		},
		{
			name: "",
			args: args{
				pais: map[pai]int{
					二筒: 2,
					三筒: 2,
					五筒: 3,
					六筒: 1,
					七筒: 1,
					八筒: 1,
					九筒: 2,
				},
			},
			want: [][]mentsu{{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HoluMianzi(tt.args.pais, tt.args.checkPai)
			for i, want := range tt.want {
				assert.ElementsMatch(t, got[i], want)
			}
		})
	}
}

func TestHuleMianziAll(t *testing.T) {
	type args struct {
		holuPattern HoluPattern
	}
	tests := []struct {
		name string
		args args
		want [][]mentsu
	}{
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 3,
						二萬: 3,
						三萬: 3,
						七索: 1,
						八索: 1,
						九索: 1,
					},
					TsumoPai: &九索,
				},
			},
			want: [][]mentsu{
				{
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{七索, 八索, 九索}},
				},
				{
					{pais: []pai{三萬, 三萬, 三萬}},
					{pais: []pai{二萬, 二萬, 二萬}},
					{pais: []pai{一萬, 一萬, 一萬}},
					{pais: []pai{七索, 八索, 九索}},
				},
			},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 1,
						二萬: 3,
						三萬: 3,
						七索: 1,
						八索: 1,
						九索: 3,
					},
					TsumoPai: &九索,
				},
			},
			want: [][]mentsu{},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 3,
						二萬: 1,
						三萬: 3,
						七索: 1,
						八索: 1,
						九索: 3,
					},
					TsumoPai: &九索,
				},
			},
			want: [][]mentsu{},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 3,
						二萬: 3,
						三萬: 1,
						七索: 1,
						八索: 1,
						九索: 3,
					},
					TsumoPai: &九索,
				},
			},
			want: [][]mentsu{},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 2,
						二萬: 2,
						三萬: 3,
						七索: 1,
						八索: 1,
						九索: 1,
						白:  2,
					},
					TsumoPai: &九索,
				},
			},
			want: [][]mentsu{},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 2,
						二萬: 2,
						三萬: 2,
						七索: 1,
						八索: 1,
						九索: 1,
						白:  3,
					},
					TsumoPai: &九索,
				},
			},
			want: [][]mentsu{
				{
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{一萬, 二萬, 三萬}},
					{pais: []pai{七索, 八索, 九索}},
					{pais: []pai{白, 白, 白}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HuleMianziAll(tt.args.holuPattern), "HuleMianziAll(%v)", tt.args.holuPattern)
		})
	}
}

func TestHoluYiban(t *testing.T) {
	type args struct {
		holuPattern HoluPattern
		rongpai     *pai
	}
	tests := []struct {
		name string
		args args
		want []StandardHoluPattern
	}{
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 3,
						二萬: 3,
						三萬: 3,
						七索: 1,
						八索: 1,
						九索: 3,
					},
					TsumoPai: &九索,
				},
			},
			want: []StandardHoluPattern{
				{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七索, 八索, 九索}},
					},
					Head:    九索,
					HoluPai: 九索,
				},
				{
					Mentsu: []mentsu{
						{pais: []pai{三萬, 三萬, 三萬}},
						{pais: []pai{二萬, 二萬, 二萬}},
						{pais: []pai{一萬, 一萬, 一萬}},
						{pais: []pai{七索, 八索, 九索}},
					},
					Head:    九索,
					HoluPai: 九索,
				},
			},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						一萬: 2,
						二萬: 2,
						三萬: 2,
						七索: 1,
						八索: 1,
						九索: 3,
						白:  3,
					},
					TsumoPai: &九索,
				},
			},
			want: []StandardHoluPattern{
				{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七索, 八索, 九索}},
						{pais: []pai{白, 白, 白}},
					},
					Head:    九索,
					HoluPai: 九索,
				},
			},
		},
		{
			name: "",
			args: args{
				holuPattern: HoluPattern{
					Menzen: map[pai]int{
						二萬: 2,
						三萬: 2,
						四萬: 2,
						五萬: 2,
						一筒: 1,
						二筒: 1,
						三筒: 1,
						二索: 1,
						三索: 1,
						四索: 1,
					},
					TsumoPai: &三萬,
				},
			},
			want: []StandardHoluPattern{
				{
					Mentsu: []mentsu{
						{pais: []pai{三萬, 四萬, 五萬}},
						{pais: []pai{三萬, 四萬, 五萬}},
						{pais: []pai{一筒, 二筒, 三筒}},
						{pais: []pai{二索, 三索, 四索}},
					},
					Head:    二萬,
					HoluPai: 三萬,
				},
				{
					Mentsu: []mentsu{
						{pais: []pai{二萬, 三萬, 四萬}},
						{pais: []pai{二萬, 三萬, 四萬}},
						{pais: []pai{一筒, 二筒, 三筒}},
						{pais: []pai{二索, 三索, 四索}},
					},
					Head:    五萬,
					HoluPai: 三萬,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HoluYiban(tt.args.holuPattern, tt.args.rongpai)
			assert.Equalf(t, tt.want, got, "HoluYiban(%v, %v)", tt.args.holuPattern, tt.args.rongpai)
		})
	}
}
