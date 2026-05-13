/*Guia de ejercicios GO
Ejercicio 34:
Cargar una matriz A de 3x5 e imprimir el mayor elemento.
*/
package main

import "fmt"

func main() {
    var A [3][5]int
    var max int

    // Cargar la matriz
    fmt.Println("Ingrese los elementos de la matriz A (3x5):")
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    // Encontrar el mayor elemento
    max = A[0][0]
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            if A[i][j] > max {
                max = A[i][j]
            }
        }
    }

    fmt.Println("El mayor elemento de la matriz es:", max)
}
