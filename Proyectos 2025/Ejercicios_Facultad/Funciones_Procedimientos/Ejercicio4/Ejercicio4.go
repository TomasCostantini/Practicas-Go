/*
Guia de ejercicios GO.
Ejercicio 23
Realizar un programa que lea dos números A y B y que le pregunte al usuario que operación desea
realizar. Las operaciones son: 1 = sumar, 2 = restar, 3 = multiplicar, 4 = dividir, 0 = salir.
Una vez que el usuario haya ingresado la operación a realizar y los números a operar, el programa
deberá mostrar el resultado de la operación y volver a preguntar al usuario que desea hacer, mientras
que el usuario no ingrese la opción 0 = salir, el programa deberá seguir funcionando.
Cada operación deberá ser implementada con distintas funciones y deberá recibir como parámetro
los números a operar y devolver el resultado de la operación.
*/
package main

import "fmt"

func main() {
	var (
		primer_numero, segundo_numero float32
		operacion                     int
	)

	fmt.Println("Elija la operacion que desee realizar(1 sumar, 2 restar, 3 multiplicar, 4 dividir, 0 salir)")
	fmt.Scan(&operacion)

	for operacion != 0 {
		fmt.Println("Ingrese el primer numero:")
		fmt.Scan(&primer_numero)
		fmt.Println("Ingrese el segundo numero:")
		fmt.Scan(&segundo_numero)

		switch operacion {
		case 1:
			fmt.Println("El resultado es:", suma(primer_numero, segundo_numero))
		case 2:
			fmt.Println("El resultado es:", resta(primer_numero, segundo_numero))
		case 3:
			fmt.Println("El resultado es:", multiplicacion(primer_numero, segundo_numero))
		case 4:
			fmt.Println("El resultado es:", division(primer_numero, segundo_numero))
		default:
			fmt.Println("Operacion no valida")
		}
		fmt.Println("Elija la operacion que desee realizar(1 sumar, 2 restar, 3 multiplicar, 4 dividir, 0 salir)")
		fmt.Scan(&operacion)

	}

	fmt.Println("Programa finalizado")
}

func suma(primer_numero, segundo_numero float32) (resultado float32) {
	resultado = primer_numero + segundo_numero
	return resultado
}

func resta(primer_numero, segundo_numero float32) (resultado float32) {
	resultado = primer_numero - segundo_numero
	return resultado
}

func multiplicacion(primer_numero, segundo_numero float32) (resultado float32) {
	resultado = primer_numero * segundo_numero
	return resultado
}

func division(primer_numero, segundo_numero float32) (resultado float32) {
	resultado = primer_numero / segundo_numero
	return resultado
}
