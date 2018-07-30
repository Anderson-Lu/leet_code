package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "测试1",
			args: args{x: 1534236469},
			want: 0,
		},
		{
			name: "测试5",
			args: args{x:1563847412},
			want: 0,
		},
		{
			name: "测试2",
			args: args{x: 213},
			want: 312,
		},
		{
			name: "测试3",
			args: args{x: -213},
			want: -312,
		},
		{
			name: "测试4",
			args: args{x: 1000},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
