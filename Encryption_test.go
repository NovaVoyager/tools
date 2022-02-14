package tools

import "testing"

func TestMD5Encryption(t *testing.T) {
	type args struct {
		password string
		salt     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "加盐测试",
			args: args{
				password: "123456",
				salt:     "test",
			},
			want: "5a2e54ee57e5b7273b9a8fed78c1ebd8",
		},
		{
			name: "不加盐测试",
			args: args{
				password: "123456",
				salt:     "",
			},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5Encryption(tt.args.password, tt.args.salt); got != tt.want {
				t.Errorf("MD5Encryption() = %v, want %v", got, tt.want)
			}
		})
	}
}
