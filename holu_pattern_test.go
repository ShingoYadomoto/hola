package hola_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mentsu_Type(t *testing.T) {
	type fields struct {
		pais []pai
	}
	tests := []struct {
		name   string
		fields fields
		want   mentsuType
	}{
		{name: "槓子", fields: fields{pais: []pai{白, 白, 白, 白}}, want: mentsuTypeKantsu},
		{name: "刻子", fields: fields{pais: []pai{一萬, 一萬, 一萬}}, want: mentsuTypeKotsu},
		{name: "順子", fields: fields{pais: []pai{一萬, 二萬, 三萬}}, want: mentsuTypeShuntsu},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mentsu{
				pais: tt.fields.pais,
			}
			assert.Equalf(t, tt.want, m.Type(), "Type()")
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
