package merge

import (
	"fmt"
	"testing"
)

// Test cases
func TestMerge(t *testing.T) {
	type test struct {
		data   List
		answer List
	}

	tests := []test{
		{
			data:   List{{25, 30}, {2, 19}, {14, 23}, {4, 8}},
			answer: List{{2, 23}, {25, 30}},
		},
		{
			data:   List{{60, 80}, {26, 42}, {2, 12}, {10, 19}, {30, 45}},
			answer: List{{2, 19}, {26, 45}, {60, 80}},
		},
		{
			data:   List{{60, 80}, {26, 42}, {2, 12}, {10, 19}, {30, 45}, {1000, 2000}},
			answer: List{{2, 19}, {26, 45}, {60, 80}, {1000, 2000}},
		},
	}

	for i, v := range tests {
		nl := Merge(v.data)
		for iv, sv := range nl {
			if sv != v.answer[iv] {
				t.Error("Test Number:", i, "Expected", v.answer, "got:", nl)
			}
		}
	}
}

// Example for Merge
func ExampleMerge() {
	l := List{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	nl := Merge(l)
	fmt.Println(nl)
	// Output:
	// [[2 23] [25 30]]
}

// Benchmark function for Merge
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := List{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
		_ = Merge(l)
	}
}
