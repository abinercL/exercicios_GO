// Faça um algoritmo para receber um número qualquer e imprimir na tela se o número é par ou ímpar, positivo ou negativo.

package main

import "fmt"

func main() {
	var a int
	fmt.Print("digite um numero :")
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("Esse valor é Positivo", a)
	} else if a < 0 {
		fmt.Println("Esse valor é negativo ", a)
	}

	if a%2 == 0 {
		fmt.Println("Esse numero é Par")
	} else {
		fmt.Println("Esse numero é Ímpar")
	}

}
