package gematria

import (
	"reflect"
	"testing"
)

func TestMajesticGematria(t *testing.T) {
	type args struct {
		in      Gematria
		choices []ComboType
	}
	tests := []struct {
		name        string
		args        args
		wantResults uint64
	}{
		{
			name: "test one",
			args: args{
				in: FromString("can pigs fly"),
				choices: []ComboType{
					CT_MJMYEI,
				},
			},
			wantResults: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := MajesticGematria(tt.args.in, tt.args.choices...); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("MajesticGematria() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
