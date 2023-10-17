// switc
// switch sem expressão é uma maneira alternativa para expressar se / mais lógica
package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2 
	fmt.Println("write", i ,"as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}

	// 2°segundo switch usando tempo e dia da semana 
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its´s the weekend")
	default:
		fmt.Println("it´s the weekday")

	}

	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("it´s before noon")
	default:
		fmt.Println("it´s after noon")

	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i´m bool")
		case int:
			fmt.Println("i´m int")
		default:
			fmt.Printf("Don't know type %t\n", t)
		}
		
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
