package gematria

import "testing"

func TestMajestic(t *testing.T) {
	type args struct {
		in uint64
	}
	tests := []struct {
		name    string
		args    args
		wantOut bool
	}{
		{
			name: "test one",
			args: args{
				in: 36,
			},
			wantOut: true,
		},
		{
			name: "test two",
			args: args{
				in: 45,
			},
			wantOut: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := Majestic(tt.args.in, nil); gotOut != tt.wantOut {
				t.Errorf("Majestic() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestSimplify(t *testing.T) {
	type args struct {
		in uint64
	}
	tests := []struct {
		name    string
		args    args
		wantOut uint64
		wantErr bool
	}{
		{
			name: "test one",
			args: args{
				in: 36,
			},
			wantOut: 9,
			wantErr: false,
		},
		{
			name: "test two",
			args: args{
				in: 1776,
			},
			wantOut: 3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := Simplify(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Simplify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("Simplify() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
