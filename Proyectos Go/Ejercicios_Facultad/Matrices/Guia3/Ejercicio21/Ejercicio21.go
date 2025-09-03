/* Guia de ejercicios GO
Ejercicio 21:
Cargar una matriz A de Nx4 elementos, donde cada fila contiene los datos correspondientes a un libro de una
librería (la primera columna contiene el código del libro, la segunda el código del autor, la tercera el número de
ejemplares y la cuarta el precio del libro). Se pide:
a) Generar un procedimiento que reciba como parámetro un código de un autor e imprima el código de todos
los libros que éste publicó. El código de autor que se debe pasar como parámetro debe ser leído antes de
invocar al procedimiento.
b) Generar una función que devuelva el código del libro y el código del autor del libro más caro. Imprimir los
resultados.
c) Generar una función que devuelva una matriz que contenga el código de libro y el código del autor de
aquellos libros cuyos números de ejemplares sea mayor a 45 unidades.
d) Ordenar la nueva matriz generado por código de autor.
e) Imprimir la matriz generada.
*/ 

package main

import "fmt"

func main() {
	var A [][4]float64
	imprimirMatriz(A)
}

func imprimirMatriz(matriz [][4]float64) {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
