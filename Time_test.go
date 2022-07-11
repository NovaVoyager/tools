package tools

import "testing"

func TestTimeStr(t *testing.T) {
	t.Log(TimeDateFormat, "-->", TimeStr(TimeDateFormat, cstZone))
	t.Log(DateFormat, "-->", TimeStr(DateFormat, cstZone))
}

func TestSecondToDayHMS(t *testing.T) {
	type args struct {
		second int64
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 int64
		want2 int64
		want3 int64
	}{
		{
			name:  "1",
			args:  args{second: 86400},
			want:  1,
			want1: 24,
			want2: 0,
			want3: 0,
		},
		{
			name:  "2",
			args:  args{second: 90301},
			want:  1,
			want1: 25,
			want2: 5,
			want3: 1,
		},
		{
			name:  "3",
			args:  args{second: 3600},
			want:  0,
			want1: 1,
			want2: 0,
			want3: 0,
		},
		{
			name:  "4",
			args:  args{second: 10},
			want:  0,
			want1: 0,
			want2: 0,
			want3: 10,
		},
		{
			name:  "4",
			args:  args{second: 5*86400 + 7200 + 300 + 1},
			want:  5,
			want1: 122,
			want2: 5,
			want3: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := SecondToDayHMS(tt.args.second)
			if got != tt.want {
				t.Errorf("SecondToDayHMS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SecondToDayHMS() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("SecondToDayHMS() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("SecondToDayHMS() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
