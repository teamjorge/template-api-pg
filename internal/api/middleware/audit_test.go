package middleware

import "testing"

func Test_cleanRemoteAddr(t *testing.T) {
	type args struct {
		remoteAddr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test expected remoteAddr",
			args: args{remoteAddr: "127.0.0.1:43901"},
			want: "127.0.0.1",
		},
		{
			name: "test no port remoteAddr",
			args: args{remoteAddr: "127.0.0.1"},
			want: "127.0.0.1",
		},
		{
			name: "test empty remoteAddr",
			args: args{remoteAddr: ""},
			want: "",
		},
		{
			name: "test ipv6 remoteAddr #1",
			args: args{remoteAddr: "[::FFFF:C0A8:1]:80"},
			want: "[::FFFF:C0A8:1]",
		},
		{
			name: "test ipv6 remoteAddr #2",
			args: args{remoteAddr: "[::FFFF:C0A8:1%1]:80"},
			want: "[::FFFF:C0A8:1%1]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanRemoteAddr(tt.args.remoteAddr); got != tt.want {
				t.Errorf("cleanRemoteAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
