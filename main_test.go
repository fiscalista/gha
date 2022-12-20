package main

import "testing"

var (
	prettyResult string
)

func TestPrettyNice(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test pretty nice",
			want: "cool",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyNice(); got != tt.want {
				t.Errorf("PrettyNice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPrettyNice(b *testing.B) {
	var result string
	for i := 0; i <= b.N; i++ {
		result = PrettyNice()
	}
	prettyResult = result
}
