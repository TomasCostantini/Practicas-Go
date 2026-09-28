/* Guia de ejercicios GO:
Ejercicio 22:
En una empresa se utiliza una matriz A de 51x6 elementos para guardar información correspondiente a las
ventas de 50 vendedores en los 4 trimestres del año. La primera columna guarda información correspondiente
al número de vendedor, a partir de la segunda columna y hasta la quinta se guardan las ventas de ese
vendedor para cada trimestre. Se pide:
a) Calcular e imprimir el total de ventas
• de cada vendedor a lo largo del año
• de cada trimestre
• anual
b) Ordenar la matriz A en forma descendente de acuerdo al total de ventas de cada vendedor.
c) Ingresar un número de vendedor e imprimir en qué trimestre realizó la mayor y la menor venta.
Nota: La fila 51 y la columna 6 pueden ser utilizadas para los fines que crea conveniente.
*/
package main

import "fmt"

func main() {

	var A [51][6]float64
	var vendedorBuscado int
	for i := 0; i < 50; i++ {

		fmt.Printf("\nNúmero del vendedor %d: ", i+1)
		fmt.Scan(&A[i][0])

		for j := 1; j <= 4; j++ {
			fmt.Printf("Ventas trimestre %d: ", j)
			fmt.Scan(&A[i][j])
		}
	}
	for i := 0; i < 50; i++ {

		A[i][5] = 0

		for j := 1; j <= 4; j++ {
			A[i][5] += A[i][j]
		}
	}
	for j := 1; j <= 4; j++ {

		A[50][j] = 0

		for i := 0; i < 50; i++ {
			A[50][j] += A[i][j]
		}
	}
	A[50][5] = 0

	for j := 1; j <= 4; j++ {
		A[50][5] += A[50][j]
	}
	fmt.Println("\n--- TOTAL POR VENDEDOR ---")

	for i := 0; i < 50; i++ {
		fmt.Printf("Vendedor %.0f: $%.2f\n",
			A[i][0], A[i][5])
	}

	fmt.Println("\n--- TOTAL POR TRIMESTRE ---")

	for j := 1; j <= 4; j++ {
		fmt.Printf("Trimestre %d: $%.2f\n",
			j, A[50][j])
	}

	fmt.Printf("\nTOTAL ANUAL EMPRESA: $%.2f\n", A[50][5])
	for i := 0; i < 49; i++ {

		for j := 0; j < 49-i; j++ {
			if A[j][5] < A[j+1][5] {
				for k := 0; k < 6; k++ {
					aux := A[j][k]
					A[j][k] = A[j+1][k]
					A[j+1][k] = aux
				}
			}
		}
	}
	fmt.Println("\n--- VENDEDORES ORDENADOS POR VENTAS ---")
	for i := 0; i < 50; i++ {
		fmt.Printf("Vendedor %.0f - Total: $%.2f\n",
			A[i][0], A[i][5])
	}
	fmt.Print("\nIngrese número de vendedor a buscar: ")
	fmt.Scan(&vendedorBuscado)
	fila := -1
	for i := 0; i < 50; i++ {
		if int(A[i][0]) == vendedorBuscado {
			fila = i
		}
	}
	if fila != -1 {
		mayor := A[fila][1]
		menor := A[fila][1]
		trimestreMayor := 1
		trimestreMenor := 1
		for j := 2; j <= 4; j++ {
			if A[fila][j] > mayor {
				mayor = A[fila][j]
				trimestreMayor = j
			}
			if A[fila][j] < menor {
				menor = A[fila][j]
				trimestreMenor = j
			}
		}
		fmt.Printf("\nVendedor %d\n", vendedorBuscado)
		fmt.Printf("Mayor venta: trimestre %d - $%.2f\n",
			trimestreMayor, mayor)
		fmt.Printf("Menor venta: trimestre %d - $%.2f\n",
			trimestreMenor, menor)
	} else {
		fmt.Println("El vendedor no existe.")
	}
}
