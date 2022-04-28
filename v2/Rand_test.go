package tools

import "testing"

func TestGetRandRang(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "1",
			args: args{
				min: 1,
				max: 10,
			},
			want: args{
				min: 1,
				max: 10,
			},
		},
	}
	for _, tt := range tests {
		for i := 0; i < 100; i++ {
			t.Run(tt.name, func(t *testing.T) {
				got := GetRandRang(tt.args.min, tt.args.max)
				if got < tt.want.min || got > tt.want.max {
					t.Errorf("GetRandRang() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
