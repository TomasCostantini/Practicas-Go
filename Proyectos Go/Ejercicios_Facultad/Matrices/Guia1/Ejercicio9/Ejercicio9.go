/*Guia de ejercicios GO
Ejercicio 35:
Cargar una matriz A de 5x3 e imprimir el menor elemento y la posición donde se encuentra (fila y columna).
*/
package main

import "fmt"

func main() {
    var A [5][3]int
    var min int
    var minRow, minCol int

    // Cargar la matriz
    fmt.Println("Ingrese los elementos de la matriz A (5x3):")
    for i := 0; i < 5; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    // Encontrar el menor elemento
    min = A[0][0]
    for i := 0; i < 5; i++ {
        for j := 0; j < 3; j++ {
            if A[i][j] < min {
                min = A[i][j]
                minRow = i
                minCol = j
            }
        }
    }

    fmt.Println("El menor elemento de la matriz es:", min)
    fmt.Println("Se encuentra en la posición (fila, columna):", minRow, minCol)
}
