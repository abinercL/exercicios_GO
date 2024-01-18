/*
	Faça um algoritmo que leia dois valores inteiros A e B, se os valores de A e B forem iguais, deverá somar os dois valores,

caso contrário devera multiplicar A por B. Ao final de qualquer um dos cálculos deve-se atribuir o resultado a uma variável C e

imprimir seu valor na tela.
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Print("digte um numero para A: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("digite um numero paras B: ")
	fmt.Scan(&b)

	var c int

	if a == b {

		c = a + b
	} else {

		c = a * b
	}
	fmt.Println("C é igual a", c)

}
