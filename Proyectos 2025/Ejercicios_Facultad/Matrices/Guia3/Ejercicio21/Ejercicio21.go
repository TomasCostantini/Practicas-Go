/* Guia de ejercicios GO
Ejercicio 21:
Cargar una matriz A de Nx4 elementos, donde cada fila contiene los datos correspondientes a un libro de una
librería (la primera columna contiene el código del libro, la segunda el código del autor, la tercera el número de
ejemplares y la cuarta el precio del libro). Se pide:
a) Generar un procedimiento que reciba como parámetro un código de un autor, la matriz e imprima el código de todos
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
import (
	"fmt"
	"math/rand"
	"time"
)
const N = 10
func main()	{
	var A [N][4]int
	cargarMatriz(&A)
	imprimirMatriz(A)
	var codigoAutor int
	fmt.Print("Ingrese el codigo de autor a buscar: ")
	fmt.Scan(&codigoAutor)
	buscarLibrosPorAutor(codigoAutor, A)
}

func cargarMatriz(matriz *[N][4]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		matriz[i][0] = rand.Intn(100) 
		matriz[i][1] = rand.Intn(10)  
		matriz[i][2] = rand.Intn(100)  
		matriz[i][3] = rand.Intn(100)  
	}
}

func imprimirMatriz(matriz [N][4]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
} 
func buscarLibrosPorAutor(codigoAutor int, matriz [N][4]int) {
	fmt.Printf("Libros del autor %d:\n", codigoAutor)
	for i := 0; i < N; i++ {
		if matriz[i][1] == codigoAutor {
			fmt.Printf("Codigo de libro: %d\n", matriz[i][0])
		}
	}
}
