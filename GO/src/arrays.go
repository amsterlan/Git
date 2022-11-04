package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("array a:", a)
	a[4] = 100
	fmt.Println("array dps de receber 100 na quarta posição", a)
	fmt.Println("o elemento na quarta posição:", a[4])
	fmt.Println("tamanho do array a:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array b:", b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("array 2d:", twoD)
}
