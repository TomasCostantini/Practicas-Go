/* Guia de ejercicios GO:
Ejercicio 22:
Los vendedores de una empresa, a fin de mes informan todas las ventas que realizaron. Por cada venta se
conoce el número de vendedor, número de producto vendido y el importe cobrado. Se comercializan 17
productos, existen 10 vendedores y el fin del ingreso de datos se produce con un número de producto nulo
(cero). Se necesita almacenar en un matriz PRO los importes totales cobrados por cada producto / vendedor.
Ordenados en forma descendente por importe total por vendedor e imprimir la matriz ordenada.
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

	// ==========================================
	// A) TOTALES
	// ==========================================

	// Total anual de cada vendedor
	// Se guarda en la columna 5
	for i := 0; i < 50; i++ {

		A[i][5] = 0

		for j := 1; j <= 4; j++ {
			A[i][5] += A[i][j]
		}
	}

	// Total de cada trimestre
	// Se guarda en la fila 50
	for j := 1; j <= 4; j++ {

		A[50][j] = 0

		for i := 0; i < 50; i++ {
			A[50][j] += A[i][j]
		}
	}

	// Total anual de toda la empresa
	A[50][5] = 0

	for j := 1; j <= 4; j++ {
		A[50][5] += A[50][j]
	}

	// ==========================================
	// IMPRESIÓN DE TOTALES
	// ==========================================

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

	// ==========================================
	// B) ORDENAR VENDEDORES DE MAYOR A MENOR
	// ==========================================

	for i := 0; i < 49; i++ {

		for j := 0; j < 49-i; j++ {

			// Comparamos usando la columna 5,
			// donde está el total del vendedor

			if A[j][5] < A[j+1][5] {

				// Intercambiamos TODA la fila
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

	// ==========================================
	// C) BUSCAR VENDEDOR
	// ==========================================

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
