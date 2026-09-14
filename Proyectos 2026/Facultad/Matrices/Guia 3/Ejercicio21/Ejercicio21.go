/*
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

const N = 4

func main() {
	var A [N][4]int

	var codigoAutor int
	fmt.Print("Ingrese el codigo de autor a buscar: ")
	fmt.Scan(&codigoAutor)

}
func cargarMatriz(matriz *[N][4]int) {
	for i := 0; i < N; i++ {
		fmt.Println("Ingrese el codigo del libro: ")
		matriz[i][0] = fmt.Println("")
		matriz[i][1] = 
		matriz[i][2] = 
		matriz[i][3] =  
	}
}

