/* Guia de ejercicios GO.
Ejercicio 21:
Una empresa de viajes que realizó una campaña de promoción necesita calcular las comisiones que
deberá pagar a sus promotores de ventas. Las mismas se calculan según la cantidad de excursiones
vendidas por cada uno de ellos. Se decide asignar una categoría a cada vendedor según la cantidad
vendida y según esa categoría se pagará un importe por cada excursión de acuerdo a la siguiente tabla:

Categoría      Excursiones         $ x Excursión
A              <10                 $1000
B              >= 10 y < 100       $1500
C              >= 100              $1900

Se ingresan las cantidades de excursiones vendidas de 10 promotores. Se necesita conocer:
a) la cantidad de promotores de cada categoría
b) el importe a pagarle a cada promotor
Resolver:
 Para el punto a) crear una función “categoriaPromotor” que recibirá como parámetro la
cantidad de excursiones que vendió y devuelva como resultado la letra de la categoría a la que
pertenece.
 Para el punto b) crear una función “importeAPagar” que recibirá como parámetro la cantidad
de excursiones que vendió, esta función deberá hacer uso de la función “categoriaPromotor”
para determinar el importe a pagar.
Imprimir los resultados.
*/

package main

import "fmt"

var (
	excursiones, inicio, monto int
	cA                         int = 0
	cB                         int = 0
	cC                         int = 0
	cantidad_promotores        int = 10
)

func main() {
	for inicio = 0; inicio < cantidad_promotores; inicio++ {
		fmt.Println("Ingrese la cantidad de excursiones vendidas:")
		fmt.Scan(&excursiones)

		fmt.Print("Promotor de categoria:", categoriaPromotor(excursiones))
		fmt.Println(" le corresponde un importe de:$", monto)
	}
	fmt.Println("Cantidad de promotores de categoria A:", cA)
	fmt.Println("Cantidad de promotores de categoria B:", cB)
	fmt.Println("Cantidad de promotores de categoria C:", cC)
}

func categoriaPromotor(excursiones int) (categoria string) {
	switch {
	case excursiones < 10:
		categoria = "A"
		cA++
	case excursiones >= 10 && excursiones < 100:
		categoria = "B"
		cB++
	case excursiones >= 100:
		categoria = "C"
		cC++
	}
	importePagar(excursiones, categoria)

	return categoria
}

func importePagar(excursiones int, categoria string) int {
	switch categoria {
	case "A":
		monto = excursiones * 1000
	case "B":
		monto = excursiones * 1500
	case "C":
		monto = excursiones * 1900
	}
	return monto
}
