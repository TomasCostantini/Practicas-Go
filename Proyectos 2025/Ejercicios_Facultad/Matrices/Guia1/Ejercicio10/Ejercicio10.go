/*Guia de ejercicios Go
Ejercicio 10:
10) Cargar una matriz A de 4x4:
a) Crear una función que calcule y devuelva como resultado la suma de los elementos de la diagonal
principal
b) Crear una función que calcule y devuelva como resultado la suma de los elementos de la triangular
superior.
c) Imprimir los resultados.
*/

package main

import "fmt"

func main() {
    var A [4][4]int
    var sumaDiagonal, sumaTriangular int

    // Cargar la matriz
    fmt.Println("Ingrese los elementos de la matriz A (4x4):")
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    // Calcular la suma de la diagonal principal
    sumaDiagonal = sumaDiagonalPrincipal(A)

    // Calcular la suma de la triangular superior
    sumaTriangular = sumaTriangularSuperior(A)

    // Imprimir los resultados
    fmt.Println("La suma de los elementos de la diagonal principal es:", sumaDiagonal)
    fmt.Println("La suma de los elementos de la triangular superior es:", sumaTriangular)
}

func sumaDiagonalPrincipal(matriz [4][4]int) int {
    suma := 0
    for i := 0; i < 4; i++ {
        suma += matriz[i][i]
    }
    return suma
}

func sumaTriangularSuperior(matriz [4][4]int) int {
    suma := 0
    for i := 0; i < 4; i++ {
        for j := i + 1; j < 4; j++ {
            suma += matriz[i][j]
        }
    }
    return suma
}
