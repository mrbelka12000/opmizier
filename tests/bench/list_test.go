package bench

import (
	"testing"
)

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
