package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repated := Repeat("a")
	expected := "aaaaa"

	if repated != expected {
		t.Errorf("expected %q but got %q", expected, repated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := Repeat("b")
	fmt.Println(repeated)
	//Output: bbbbb
}
