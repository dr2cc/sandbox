package main

import "testing"

func TestIsEven(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple Even ", args{4}, "Even"},
		{"simple Odd ", args{0}, "Even"},
		{"-1 - нечетное число", args{-1}, "Odd"},
		{"0 - четное число", args{0}, "Even"},
		{"1 - нечетное число", args{1}, "Odd"},
		{"2 - четное число", args{2}, "Even"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.input); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
