/*Guia de ejercicios GO
Ejercicio 33:
Cargar una matriz A de 5x5 y contar aquellos elementos que coinciden con su fila o columna.
*/
package main

import "fmt"

func main() {
    var A [5][5]int
    var count int

    // Cargar la matriz
    fmt.Println("Ingrese los elementos de la matriz A (5x5):")
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    // Contar elementos que coinciden con su fila o columna
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            if A[i][j] == i || A[i][j] == j {
                count++
            }
        }
    }

    fmt.Println("Cantidad de elementos que coinciden con su fila o columna:", count)
}