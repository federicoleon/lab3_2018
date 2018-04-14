package main

import (
	ordenamiento "github.com/federicoleon/lab3_2018/sort"
	"log"
)

func main() {
	palin := ordenamiento.EsPalindromo("anita lava la tina", true)
	log.Printf("Es palindromo %v", palin)
}
