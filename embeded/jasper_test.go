package embeded

import (
	"testing"
)

/*
Tested on Linux
*/
// func TestMainC(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{"Test MainC"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			MainC()
// 		})
// 	}
// 	t.Error("Done.")
// }

func TestMainMethodC(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMainMethodC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			MainMethodC()
		})
	}
}

func TestSquareMethodC(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSquareMethodC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			SquareMethodC()
		})
	}
}

func TestHelloWorld(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestHelloWorld", args{"Chatchai Vichai1"}, "Chatchai Vichai1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HelloWorld(tt.args.name)
			t.Log(got)
			if got != tt.want {
				t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestALL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSquareMethodC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			SquareMethodC()
			got := HelloWorld("Chatchai Vichai")
			t.Log(got)
			MainMethodC()
			PowerC()
		})
	}
}
