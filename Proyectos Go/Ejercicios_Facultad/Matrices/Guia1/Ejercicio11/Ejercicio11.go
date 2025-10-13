/*Guia de ejercicios Go
Ejercicio 11:
Cargar una matriz A de 6x6, generar una nueva matriz de igual dimensión a la dada bajo las siguientes
condiciones:
a) si el elemento es negativo colocar un 0 en la nueva matriz.
b) si el elemento es cero colocar un 1 en la nueva matriz.
c) si el elemento es positivo colocar un 2 en la nueva matriz.
d) Generar un procedimiento que imprima la matriz pasada por parámetro. Utilizar el procedimiento para
imprimir las dos matrices
*/
package main

import "fmt"

func main() {
    var A [6][6]int
    var B [6][6]int

    // Cargar la matriz A
    fmt.Println("Ingrese los elementos de la matriz A (6x6):")
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    // Generar la matriz B
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            if A[i][j] < 0 {
                B[i][j] = 0
            } else if A[i][j] == 0 {
                B[i][j] = 1
            } else {
                B[i][j] = 2
            }
        }
    }

    // Imprimir las matrices
    fmt.Println("Matriz A:")
    imprimirMatriz(A)
    fmt.Println("Matriz B:")
    imprimirMatriz(B)
}

func imprimirMatriz(matriz [6][6]int) {
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            fmt.Print(matriz[i][j], " ")
        }
        fmt.Println()
    }
}
