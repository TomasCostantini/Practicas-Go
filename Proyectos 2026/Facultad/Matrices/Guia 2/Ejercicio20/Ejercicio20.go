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

import "fmt"

const (
	filas      int     = 4
	columnas   int     = 7
	materias   int     = 6
	col        int     = 8
	notaMinima float32 = 60
)

func main() {
	var (
		A [filas][columnas]float32
		C [filas][col]float32
	)

	for f := 0; f < filas; f++ {
		fmt.Print("Ingrese el numero de legajo:")
		fmt.Scanln(&A[f][0])
		for c := 1; c < columnas; c++ {
			fmt.Printf("Ingrese la nota de la %v° materia:", c)
			fmt.Scanln(&A[f][c])
		}
	}
	M := generarMatrizM(A)
	C = generarMatrizC(A)
	ordenarMatrizC(&C)
	fmt.Println("La matriz de los que que Poseen un numero superior o igual a 60 en todas las materias")
	for f := 0; f < len(M); f++ {
		for c := 0; c < columnas; c++ {
			fmt.Printf("|%4v |", M[f][c])
		}
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("La matriz ordenada en forma ascendete segun los promedios")
	for f := 0; f < filas; f++ {
		for c := 0; c < col; c++ {
			fmt.Printf("|%4v |", C[f][c])
		}
		fmt.Println("")
	}
}

func generarMatrizM(A [filas][columnas]float32) (M [][columnas]float32) {
	var (
		mat    int
		vector [columnas]float32
	)
	for f := 0; f < filas; f++ {
		mat = 0
		for c := 1; c < columnas; c++ {
			if A[f][c] >= notaMinima {
				mat++
			}
			if mat == materias {
				vector[0] = A[f][0]
				for c := 1; c < columnas; c++ {
					vector[c] = A[f][c]
				}
				M = append(M, vector)
			}
		}
	}

	return M
}
func generarMatrizC(A [filas][columnas]float32) (C [filas][col]float32) {
	var (
		suma float32
	)
	for f := 0; f < filas; f++ {
		C[f][0] = A[f][0]
		suma = 0
		for c := 1; c < columnas; c++ {
			suma += A[f][c]
			C[f][c] = A[f][c]
		}
		promedio := suma / float32(materias)
		C[f][columnas] = promedio
	}

	return C
}

func ordenarMatrizC(C *[filas][col]float32) {

	for f1 := 0; f1 < filas-1; f1++ {
		for f2 := f1 + 1; f2 < filas; f2++ {
			if C[f1][7] > C[f2][7] {
				for c := 0; c < col; c++ {
					aux := C[f1][c]
					C[f1][c] = C[f2][c]
					C[f2][c] = aux
				}
			}
		}
	}
}