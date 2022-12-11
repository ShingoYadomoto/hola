package hola_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHupaiCalculater_ipeko(t *testing.T) {
	type fields struct {
		standard StandardHoluPattern
	}
	tests := []struct {
		name   string
		fields fields
		want   []HandType
	}{
		{
			name: "true case. ipeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{一盃口},
		},
		{
			name: "false case. ipeko case. 3 same shuntsu",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
					},
				},
			},
			want: []HandType{一盃口},
		},
		{
			name: "false case. ipeko case. 3 same shuntsu",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{東, 東, 東}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. ryanpeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. fulou case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
					},
					FulouMentsu: FulouMentsuList{
						{mentsu: mentsu{pais: []pai{白, 白, 白}}},
					},
				},
			},
			want: []HandType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := HupaiCalculater{
				standard: tt.fields.standard,
			}
			assert.Equalf(t, tt.want, hc.ipeko(), "ipeko()")
		})
	}
}

func TestHupaiCalculater_ryanpeko(t *testing.T) {
	type fields struct {
		standard   StandardHoluPattern
		zhuangfeng Zhuangfeng
		zifeng     Zifeng
	}
	tests := []struct {
		name   string
		fields fields
		want   []HandType
	}{
		{
			name: "true case. ryanpeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{七萬, 八萬, 九萬}},
					},
				},
			},
			want: []HandType{二盃口},
		},
		{
			name: "false case. ipeko case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
		{
			name: "false case. not same shuntsu case",
			fields: fields{
				standard: StandardHoluPattern{
					Mentsu: []mentsu{
						{pais: []pai{一萬, 二萬, 三萬}},
						{pais: []pai{東, 東, 東}},
						{pais: []pai{中, 中, 中}},
						{pais: []pai{白, 白, 白}},
					},
				},
			},
			want: []HandType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := HupaiCalculater{
				standard:   tt.fields.standard,
				zhuangfeng: tt.fields.zhuangfeng,
				zifeng:     tt.fields.zifeng,
			}
			assert.Equalf(t, tt.want, hc.ryanpeko(), "ryanpeko()")
		})
	}
}
