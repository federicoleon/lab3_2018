// This is the sort package
package sort

import (
	"strings"
)

func Sort(list []int) {
	sigo := true
	for sigo {
		sigo = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sigo = true
			}
		}
	}
}

func EsPalindromo(s string) bool {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", len(s))

	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return EsPalindromo(s[1 : len(s)-1])
}
