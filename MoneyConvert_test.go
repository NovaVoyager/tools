package tools

import "testing"

func TestFen2Yuan(t *testing.T) {
	type args struct {
		fen int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "100分转1元",
			args: args{fen: 100},
			want: 1.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fen2Yuan(tt.args.fen); got != tt.want {
				t.Errorf("Fen2Yuan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYuan2Fen(t *testing.T) {
	type args struct {
		yuan float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1元转100分",
			args: args{yuan: 1},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Yuan2Fen(tt.args.yuan); got != tt.want {
				t.Errorf("Yuan2Fen() = %v, want %v", got, tt.want)
			}
		})
	}
}
