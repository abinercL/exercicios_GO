// Faça um algoritmo que leia os valores de A, B, C e em seguida imprima na tela a soma entre A e B é mostre se a soma é menor que C.
package main

import "fmt"

func main() {

	var a int
	fmt.Print("digite um numero para A: ")
	fmt.Scan(&a)
	fmt.Println("o valor é", a)

	var b int
	fmt.Print("digite um numero para B: ")
	fmt.Scan(&b)
	fmt.Println("o valor é", b)

	var c int
	fmt.Print("digite um numero para C: ")
	fmt.Scan(&c)
	fmt.Println("o valor é", c)

	num := a + b

	if num == c {
		fmt.Println(num, "é igual a", c)
	} else if num < c {
		fmt.Println(num, "é menor que", c)
	} else {
		fmt.Println(num, "é maior que ", c)
	}

}
