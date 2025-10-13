/*
Guia de ejercicios GO:
Ejercicio 20:
Una casa de artículos deportivos almacena en una matriz A de 100x4 los datos correspondientes a los 100
artículos que vende, donde la primera columna contiene el código del artículo, la segunda el código de la
marca del artículo, la tercera la cantidad existente y la cuarta el precio de venta. Se pide:
a) Generar un procedimiento que reciba como parámetro la matriz A e imprimir el código de artículo y el
precio correspondiente de aquellos cuya cantidad existente supere las cincuenta unidades.
b) Generar una función que reciba como parámetro la matriz A y devuelva como resultado una matriz con el
código del artículo y el código de marca de aquellos artículos cuyo precio sea menor o igual a $300.
c) Generar una función que ordene en forma ascendente la matriz A de acuerdo al precio de venta. La
función deberá recibir como parámetro por referencia la matriz a ordenar y el número de columna de la
matriz por la que se desea ordenar.
d) Generar una función que imprima la matriz recibida como parámetro. Utilizarla para imprimir la matriz A.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100
const M = 4

func main() {
	var A [N][M]int
	cargarMatriz(&A)
	imprimirMatriz(A)
	articulosConCantidadMayorACincuenta(A)
	resultado := articulosConPrecioMenorOIgualA300(A)
	imprimirMatrizGenerica(resultado)
	ordenarPorPrecio(&A)
	imprimirMatriz(A)
}
func cargarMatriz(matriz *[N][M]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		matriz[i][0] = rand.Intn(100) 
		matriz[i][1] = rand.Intn(10)  
		matriz[i][2] = rand.Intn(100) 
		matriz[i][3] = rand.Intn(100) 
	}
}
func imprimirMatriz(matriz [N][M]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
func articulosConCantidadMayorACincuenta(matriz [N][M]int) {
	fmt.Println("Articulos con cantidad mayor a 50:")
	for i := 0; i < N; i++ {
		if matriz[i][2] > 50 {
			fmt.Printf("Codigo de articulo: %d, Precio: %d\n", matriz[i][0], matriz[i][3])
		}
	}
}
func articulosConPrecioMenorOIgualA300(matriz [N][M]int) [][2]int {
	var resultado [][2]int
	for i := 0; i < N; i++ {
		if matriz[i][3] <= 300 {
			resultado = append(resultado, [2]int{matriz[i][0], matriz[i][1]})
		}
	}
	return resultado
}
func ordenarPorPrecio(matriz *[N][M]int) {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if matriz[i][3] > matriz[j][3] {
				for k := 0; k < M; k++ {
					matriz[i][k], matriz[j][k] = matriz[j][k], matriz[i][k]
				}
			}
		}
	}
}
func imprimirMatrizGenerica(matriz [][2]int) {
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
