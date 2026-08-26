/*
Guia de ejercicios GO:
Ejercicio 2:
Cargar una matriz A de 6x5, Crear una función que reciba como parámetro la matriz cargada y devuelva la suma
de aquellos elementos que sean mayor o igual que 5 y la cantidad de elementos que intervinieron. Imprimir los
resultados.
*/
package main

import "fmt"

const (
	filas    = 6
	columnas = 5
)

func main() {
	var (
		A      [filas][columnas]int
		numero int
	)
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			fmt.Println("Ingrese un numero: ")
			fmt.Scan(&numero)
			A[i][z] = numero
		}
	}
	suma, elementos := suma(A)
	fmt.Println("El valor de la suma es: ", suma)
	fmt.Println("El numero de elementos sumados es: ", elementos)
	fmt.Println("Progrma finalizado")
}

func suma(A [filas][columnas]int) (int, int) {
	var (
		suma      = 0
		elementos = 0
	)
	for i := 0; i < filas; i++ {
		for z := 0; z < columnas; z++ {
			if A[i][z] >= 5 {
				suma = suma + A[i][z]
				elementos = elementos + 1
			}
		}
	}
	return suma, elementos
}
