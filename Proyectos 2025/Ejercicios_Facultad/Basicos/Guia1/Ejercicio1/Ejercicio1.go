/* Guia de ejercicios Go.
Ejercicio 1:
.Leer tres números enteros y realizar:
• la diferencia del primero respecto del segundo. Imprimir el resultado.
• el productos de los dos últimos. Imprimir el resultado.
• la división entera entre el primero y el tercero. Imprimir el resultado.*/

package main

import "fmt"

func main() {
	var (
		primer_valor  int
		segundo_valor int
		tercer_valor  int
		diferencia    int
		producto      int
		division      int
	)

	println("Ingrese el primer valor: ")
	fmt.Scanln(&primer_valor)
	println("Ingrese el segundo valor: ")
	fmt.Scanln(&segundo_valor)
	println("Ingrese el tercer valor: ")
	fmt.Scanln(&tercer_valor)

	diferencia = primer_valor - segundo_valor
	producto = segundo_valor + tercer_valor
	division = primer_valor / tercer_valor

	println("La diferencia es:", diferencia)
	println("El producto es:", producto)
	println("La division es:", division)

}
