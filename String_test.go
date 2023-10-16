package tools

import "testing"

func TestKrand(t *testing.T) {
	t.Log(KcRandKindNum, KRand(6, KcRandKindNum))
	t.Log(KcRandKindLower, KRand(6, KcRandKindLower))
	t.Log(KcRandKindUpper, KRand(6, KcRandKindUpper))
	t.Log(KcRandKindAll, KRand(6, KcRandKindAll))
}

func TestFirstUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "tom",
			},
			want: "Tom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUpper(tt.args.s); got != tt.want {
				t.Errorf("FirstUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrEscape(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{str: `'a'a b'bbb' "c"c "d"d`},
			want: `\'a\'a b\'bbb\' \"c\"c \"d\"d`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrEscape(tt.args.str); got != tt.want {
				t.Errorf("StrEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
