/* Guia de ejercicios GO:
Ejercicio 23:
Los vendedores de una empresa, a fin de mes informan todas las ventas que realizaron. Por cada venta se
conoce el número de vendedor, número de producto vendido y el importe cobrado. Se comercializan 17
productos, existen 10 vendedores y el fin del ingreso de datos se produce con un número de producto nulo
(cero). Se necesita almacenar en un matriz PRO los importes totales cobrados por cada producto / vendedor.
Ordenados en forma descendente por importe total por vendedor e imprimir la matriz ordenada.
*/

package main

import "fmt"

func main() {

	var PRO [10][17]float64
	var vendedor, producto int
	var importe float64

	fmt.Print("Ingrese número de producto (0 para finalizar): ")
	fmt.Scan(&producto)

	for producto != 0 {

		fmt.Print("Ingrese número de vendedor (1-10): ")
		fmt.Scan(&vendedor)

		fmt.Print("Ingrese importe de la venta: ")
		fmt.Scan(&importe)

		PRO[vendedor-1][producto-1] += importe

		fmt.Print("\nIngrese número de producto (0 para finalizar): ")
		fmt.Scan(&producto)
	}
	var totalVendedor [10]float64

	for i := 0; i < 10; i++ {
		for j := 0; j < 17; j++ {
			totalVendedor[i] += PRO[i][j]
		}
	}
	var numVendedor [10]int
	for i := 0; i < 10; i++ {
		numVendedor[i] = i + 1
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9-i; j++ {

			if totalVendedor[j] < totalVendedor[j+1] {
				totalVendedor[j], totalVendedor[j+1] =
					totalVendedor[j+1], totalVendedor[j]
				numVendedor[j], numVendedor[j+1] =
					numVendedor[j+1], numVendedor[j]
			}
		}
	}
	fmt.Println("\n--- VENDEDORES ORDENADOS POR IMPORTE TOTAL ---")

	for i := 0; i < 10; i++ {
		fmt.Printf("Vendedor %d: $%.2f\n",
			numVendedor[i], totalVendedor[i])
	}
}
