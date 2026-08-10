/*
Guia de ejercicios GO:
Ejercicio 3:
Una empresa de viajes que realizó una campaña de promoción necesita calcular las comisiones que
deberá pagar a sus promotores de ventas. Las mismas se calculan según la cantidad de excursiones
vendidas por cada uno de ellos. Se decide asignar una categoría a cada vendedor según la cantidad
vendida y según esa categoría se pagará un importe por cada excursión de acuerdo a la siguiente tabla:

Categoría         Excursiones            $ x Excursión
A                   <10                     $1000
B               >= 10 y < 100               $1500
C                  >= 100                   $1900

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

const (
	promotores = 10
	importeA   = 1000
	importeB   = 1500
	importeC   = 1900
)

var (
	excursiones, comision int
	promotoresA = 0
	promotoresB = 0
	PromotoresC = 0
	categoria   string
)

func main() {
	for i := 0; i < promotores; i++ {
		fmt.Println("Ingrese la cantidad de excursiones vendidas: ")
		fmt.Scan(&excursiones)
		fmt.Println("El promotor es de categoria: ", categoriaPromotor(excursiones))
		fmt.Println("La comision correspondiente es: ",importePagar(categoria,importeA,importeB,importeC))
	}
}

func categoriaPromotor(excursiones int) string {
	switch {
	case excursiones < 10:
		categoria = "A"
		promotoresA++
	case excursiones > 10 && excursiones < 100:
		categoria = "B"
		promotoresB++
	case excursiones > 100:
		categoria = "C"
		PromotoresC++
	}
	return categoria
}

func importePagar(categoria,importeA,importeB,importeC int) int{

	return comision
}
