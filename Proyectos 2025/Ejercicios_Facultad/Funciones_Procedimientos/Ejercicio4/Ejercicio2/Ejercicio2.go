/*
Guia de ejercicios GO:
Ejercicio 21:
Leer 5 números e indicar por cada uno si es múltiplos de 3 y no de 5. Para ello crear una función que se
llame “verificarMultiplo” que reciba 2 parámetro:
 El número que se desea verificar si es múltiplo de.
 El múltiplo a verificar (3 o 5).
La función deberá devolver como resultado un valor booleano indicando si es múltiplo o no. Para
verificar si la multiplicidad se cumple se deberá realizar la división restando.
El programa a desarrollar deberá implementar la función para resolver el problema planteado.
*/
package main

import "fmt"

func main() {
	var (
		inicio_secuencia              int
		final_secuencia               int = 5
		primer_numero, primer_numerop int
	)
	for inicio_secuencia = 0; inicio_secuencia < final_secuencia; inicio_secuencia++ {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&primer_numero)
		primer_numerop = primer_numero

		if verificarMultiplo(primer_numero, primer_numerop) == true {
			fmt.Println("El numero cumple con la condicion")
		} else {
			fmt.Println("El numero no cumple con la condicion")
		}

	}
}

func verificarMultiplo(primer_numero, primer_numerop int) (resultado bool) {
	for primer_numero >= 3 {
		primer_numero -= 3
	}
	for primer_numerop >= 5 {
		primer_numerop -= 5
	}
	if primer_numero == 0 && primer_numerop != 0 {
		resultado = true
	} else {
		resultado = false
	}
	return resultado
}
