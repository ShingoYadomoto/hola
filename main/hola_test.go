package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHolaGuoshi(t *testing.T) {
	type args struct {
		hand    Hand
		rongpai *pai
	}
	tests := []struct {
		name string
		args args
		want *KokushiHolaHand
	}{
		{
			name: "tsumo",
			args: args{
				hand: Hand{
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
			want: &KokushiHolaHand{
				Head:    中,
				HolaPai: 發,
			},
		},

		{
			name: "tsumo. 13面",
			args: args{
				hand: Hand{
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
			want: &KokushiHolaHand{
				Head:    中,
				HolaPai: 中,
			},
		},

		{
			name: "rong",
			args: args{
				hand: Hand{
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
			want: &KokushiHolaHand{
				Head:    中,
				HolaPai: 發,
			},
		},

		{
			name: "rong. 13面",
			args: args{
				hand: Hand{
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
			want: &KokushiHolaHand{
				Head:    中,
				HolaPai: 中,
			},
		},

		{
			name: "not kokushi",
			args: args{
				hand: Hand{
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
			if got := HolaGuoshi(tt.args.hand, tt.args.rongpai); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HolaGuoshi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHolaGiduizi(t *testing.T) {
	type args struct {
		hand    Hand
		rongpai *pai
	}
	tests := []struct {
		name string
		args args
		want *TitoitsuHolaHand
	}{
		{
			name: "tsumo",
			args: args{
				hand: Hand{
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
			want: &TitoitsuHolaHand{
				Hand:    []pai{一萬, 九萬, 一筒, 九筒, 一索, 九索, 東},
				HolaPai: 發,
			},
		},

		{
			name: "rong",
			args: args{
				hand: Hand{
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
			want: &TitoitsuHolaHand{
				Hand:    []pai{一萬, 九萬, 一筒, 九筒, 一索, 九索, 東},
				HolaPai: 發,
			},
		},

		{
			name: "not Titoitsu",
			args: args{
				hand: Hand{
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
			got := HolaGiduizi(tt.args.hand, tt.args.rongpai)
			if tt.want != nil {
				assert.ElementsMatch(t, got.Hand, tt.want.Hand)
				assert.Equal(t, got.HolaPai, tt.want.HolaPai)
			} else {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestHolaMianzi(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HolaMianzi(tt.args.pais, tt.args.checkPai)
			for i, want := range tt.want {
				assert.ElementsMatch(t, got[i], want)
			}
		})
	}
}
