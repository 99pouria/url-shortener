package converter

import "testing"

func BenchmarkItoA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ItoA(1)
	}
}
