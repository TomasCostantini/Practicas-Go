/* Guia de ejercicios Go.
Ejercicio 19:
Una empresa de viajes que realizó una campaña de promoción necesita calcular las comisiones que deberá
pagar a sus promotores de ventas. Las mismas se calculan según la cantidad de excursiones vendidas por
cada uno de ellos. Se decide asignar una categoría a cada vendedor según la cantidad vendida y según esa
categoría se pagará un importe por cada excursión de acuerdo a la siguiente tabla:

Categoría         Excursiones             $ x Excursión
A                  <10                    $1000
B                  >= 10 y < 50           $1500
C                  >= 50 y < 100          $1700
D                  >= 100                 $1900

Se cuenta con N datos que corresponden a las cantidades vendidas por cada promotor. Se necesita
conocer:
• El importe a pagarle a cada promotor
• La cantidad de promotores de cada categoría
*/

package main

import "fmt"

func main() {
	var (
		cantidadpromotores int
		excursiones        int
		promotoresA        int = 0
		promotoresB        int = 0
		promotoresC        int = 0
		promotoresD        int = 0
		valor              int
	)

	fmt.Println("Ingrese la cantidad de promotores: ")
	fmt.Scan(&cantidadpromotores)

	for i := 0; i < cantidadpromotores; i++ {
		fmt.Println("Ingrese la cantidad de excursiones vendidas :")
		fmt.Scan(&excursiones)

		if excursiones < 10 {
			promotoresA++
			valor = excursiones * 1000
			fmt.Println("Este promotor es de categoria A")
			fmt.Println("El valor de las comisiones para este promotor es de: ", valor)
		}

		if excursiones >= 10 && excursiones < 50 {
			promotoresB++
			valor = excursiones * 1500
			fmt.Println("Este promotor es de categoria B")
			fmt.Println("El valor de las comisiones para este promotor es de: ", valor)
		}

		if excursiones >= 50 && excursiones < 100 {
			promotoresC++
			valor = excursiones * 1700
			fmt.Println("Este promotor es de categoria C")
			fmt.Println("El valor de las comisiones para este promotor es de: ", valor)
		}
		if excursiones > 100 {
			promotoresD++
			valor = excursiones * 1900
			fmt.Println("Este promotor es de categoria D")
			fmt.Println("El valor de las comisiones para este promotor es de: ", valor)
		}
	}

	fmt.Println("Cantidad de promotores por categoria: ")
	fmt.Println("Categoria A: ", promotoresA)
	fmt.Println("Categoria B: ", promotoresB)
	fmt.Println("Categoria C: ", promotoresC)
	fmt.Println("Categoria D: ", promotoresD)

}
