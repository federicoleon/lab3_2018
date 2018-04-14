package sort

import (
	"testing"
)

func TestIsPalindromo(t *testing.T) {
	a := "anita lava la tina"
	palindromo := EsPalindromo(a)

	if !palindromo {
		t.Error("anita lava la tina debería ser palindromo")
	}
}

func BenchmarkBurbuja(b *testing.B) {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := 0; i < b.N; i++ {
		Sort(a)
	}
}

func BenchmarkPalindromo(b *testing.B) {
	a := "anitalavalatina"

	for i := 0; i < b.N; i++ {
		EsPalindromo(a)
	}
}
