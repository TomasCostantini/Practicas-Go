/*Guia de ejercicios Go
Ejercicio 13:
Cargar una matriz A de 10x20 elementos. Generar una función que devuelva como resultado el mayor
elemento de las filas pares y el menor de las filas impares y las posiciones donde se ubican. La función deberá
recibir como parámetros:
a) La matriz a buscar
b) Si se desea buscar el mayor o menor (parámetro booleano)
c) Si se desea buscar sobre filas pares o impares (parámetro booleano)
Imprimir los resultados.
*/
package main

import "fmt"

func main() {
    var A [10][20]int
    // Cargar la matriz A
    fmt.Println("Ingrese los elementos de la matriz A (10x20):")
    for i := 0; i < 10; i++ {
        for j := 0; j < 20; j++ {
            fmt.Scan(&A[i][j])
        }
    }

    // Buscar el mayor de las filas pares y el menor de las filas impares
    mayorPar, posMayorPar := buscarElemento(A, true, true)
    menorImpar, posMenorImpar := buscarElemento(A, false, false)

    // Imprimir los resultados
    fmt.Println("Mayor elemento de filas pares:", mayorPar, "en posición:", posMayorPar)
    fmt.Println("Menor elemento de filas impares:", menorImpar, "en posición:", posMenorImpar)
}

func buscarElemento(matriz [10][20]int, buscarMayor bool, filasPares bool) (int, [2]int) {
    var resultado int
    var posicion [2]int
    if buscarMayor {
        resultado = -1
    } else {
        resultado = 1001
    }

    for i := 0; i < 10; i++ {
        if (filasPares && i%2 == 0) || (!filasPares && i%2 != 0) {
            for j := 0; j < 20; j++ {
                if (buscarMayor && matriz[i][j] > resultado) || (!buscarMayor && matriz[i][j] < resultado) {
                    resultado = matriz[i][j]
                    posicion = [2]int{i, j}
                }
            }
        }
    }
    return resultado, posicion
}
