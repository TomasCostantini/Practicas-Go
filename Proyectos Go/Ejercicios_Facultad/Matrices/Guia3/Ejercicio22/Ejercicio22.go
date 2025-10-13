/*
	Guia de ejercicios GO

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

const (
	filas    = 51
	columnas = 6
)

func main() {
	var matrizA [filas][columnas]int
	var ventas int
	var nVendedor int

	for f := 0; f < filas-1; f++ {
		fmt.Print("Ingrese el Número de Vendedor: ")
		fmt.Scanln(&nVendedor)
		matrizA[f][0] = nVendedor

		for c := 1; c <= 4; c++ {
			fmt.Printf("Ingrese las ventas del Trimestre %d: ", c)
			fmt.Scanln(&ventas)
			matrizA[f][c] = ventas
		}
	}

	TotalVentasVendedores(&matrizA)
	TotalVentasTrimestres(&matrizA)
	TotalVentasAnual(&matrizA)

	OrdenarVendedores(&matrizA)

	ConsultaVendedor(&matrizA)
}

func TotalVentasVendedores(matrizA *[filas][columnas]int) {
	for f := 0; f < filas-1; f++ {
		acumulador := 0
		for c := 1; c <= 4; c++ {
			acumulador += matrizA[f][c]
		}
		matrizA[f][5] = acumulador
		fmt.Printf("El total anual del vendedor %d es %d\n", matrizA[f][0], acumulador)
	}
	fmt.Println()
}

func TotalVentasTrimestres(matrizA *[filas][columnas]int) {
	for c := 1; c <= 4; c++ {
		acumulador := 0
		for f := 0; f < filas-1; f++ {
			acumulador += matrizA[f][c]
		}
		matrizA[filas-1][c] = acumulador
		fmt.Printf("El total de ventas del Trimestre %d es %d\n", c, acumulador)
	}
	fmt.Println()
}

func TotalVentasAnual(matrizA *[filas][columnas]int) {
	total := 0
	for c := 1; c <= 4; c++ {
		total += matrizA[filas-1][c]
	}
	matrizA[filas-1][5] = total
	fmt.Printf("El total anual de la empresa es %d\n\n", total)
}

func OrdenarVendedores(matrizA *[filas][columnas]int) {
	var auxFila [columnas]int
	for i := 0; i < filas-2; i++ {
		for j := i + 1; j < filas-1; j++ {
			if matrizA[i][5] < matrizA[j][5] {
				auxFila = matrizA[i]
				matrizA[i] = matrizA[j]
				matrizA[j] = auxFila
			}
		}
	}

	fmt.Println("Matriz de vendedores ordenada por total anual descendente:")
	for f := 0; f < filas-1; f++ {
		fmt.Printf("%d\n", matrizA[f])
	}
	fmt.Println()
}

func ConsultaVendedor(matrizA *[filas][columnas]int) {
	var numConsulta int
	fmt.Print("Ingrese el número de vendedor a consultar: ")
	fmt.Scanln(&numConsulta)

	for f := 0; f < filas-1; f++ {
		if matrizA[f][0] == numConsulta {
			maxVenta := matrizA[f][1]
			minVenta := matrizA[f][1]
			trimestreMax := 1
			trimestreMin := 1
			for c := 2; c <= 4; c++ {
				if matrizA[f][c] > maxVenta {
					maxVenta = matrizA[f][c]
					trimestreMax = c
				}
				if matrizA[f][c] < minVenta {
					minVenta = matrizA[f][c]
					trimestreMin = c
				}
			}
			fmt.Printf("El vendedor %d vendió menos en el Trimestre %d y más en el Trimestre %d\n", numConsulta, trimestreMin, trimestreMax)
			break
		}
	}
}
