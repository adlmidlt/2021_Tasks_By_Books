package main

import "testing"

func Test_getName(t *testing.T) {

	name := getName()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Ok",
			want: name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getName(); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}
