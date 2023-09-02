package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expectedResult := strings.Repeat("a", 5)

	if result != expectedResult {
		t.Errorf("expected %q but got %q", expectedResult, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
