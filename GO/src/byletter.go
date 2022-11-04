package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"Ronaldo", "Alan", "Lucas", "Robson", "Vaneila",
		"falcon", "Romario", "Tony", "Bruna"}

	filtered := Filter(words, func(word string) bool {
		return strings.HasPrefix(word, "R")
	})

	fmt.Println(filtered)

}

func Filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
