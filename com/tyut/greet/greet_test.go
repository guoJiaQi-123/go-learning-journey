package greet

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello,Chris"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func Greet(b *bytes.Buffer, name string) {
	fmt.Fprintf(b, "Hello,%s", name)
}
