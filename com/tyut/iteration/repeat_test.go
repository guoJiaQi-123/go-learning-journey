package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 8)
	want := "aaaaaaaa"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func Repeat(s string, frequency int) string {
	var repeated string
	for i := 0; i < frequency; i++ {
		repeated += s
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 2)
	fmt.Println(repeat)
	// Output: aa
}
