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
