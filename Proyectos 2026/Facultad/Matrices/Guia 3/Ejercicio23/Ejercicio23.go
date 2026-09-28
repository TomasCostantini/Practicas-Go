package main

import "fmt"

func main() {

	var PRO [10][17]float64
	var vendedor, producto int
	var importe float64

	// CARGA DE VENTAS
	fmt.Print("Ingrese número de producto (0 para finalizar): ")
	fmt.Scan(&producto)

	for producto != 0 {

		fmt.Print("Ingrese número de vendedor (1-10): ")
		fmt.Scan(&vendedor)

		fmt.Print("Ingrese importe de la venta: ")
		fmt.Scan(&importe)

		// Restamos 1 porque los índices comienzan en 0
		PRO[vendedor-1][producto-1] += importe

		fmt.Print("\nIngrese número de producto (0 para finalizar): ")
		fmt.Scan(&producto)
	}

	// TOTAL VENDIDO POR CADA VENDEDOR
	var totalVendedor [10]float64

	for i := 0; i < 10; i++ {
		for j := 0; j < 17; j++ {
			totalVendedor[i] += PRO[i][j]
		}
	}

	// Vector para conservar el número de vendedor
	var numVendedor [10]int

	for i := 0; i < 10; i++ {
		numVendedor[i] = i + 1
	}

	// ORDENAMIENTO DESCENDENTE
	// Bubble Sort
	for i := 0; i < 9; i++ {
		for j := 0; j < 9-i; j++ {

			if totalVendedor[j] < totalVendedor[j+1] {

				// Intercambio de totales
				totalVendedor[j], totalVendedor[j+1] =
					totalVendedor[j+1], totalVendedor[j]

				// Intercambio de vendedores
				numVendedor[j], numVendedor[j+1] =
					numVendedor[j+1], numVendedor[j]
			}
		}
	}

	// IMPRESIÓN
	fmt.Println("\n--- VENDEDORES ORDENADOS POR IMPORTE TOTAL ---")

	for i := 0; i < 10; i++ {
		fmt.Printf("Vendedor %d: $%.2f\n",
			numVendedor[i], totalVendedor[i])
	}
}