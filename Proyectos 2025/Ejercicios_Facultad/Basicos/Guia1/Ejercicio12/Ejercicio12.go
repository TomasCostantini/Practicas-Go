/* Guia de ejercicios GO
Ejercicio 12:
Leer 10 números enteros y determinar cuántos son múltiplos de 5. Imprimir el resultado. (utilizar %)
*/
package main
import "fmt"

func main() {


    var (
        numero   int
        contador int
    )



    for i := 0; i < 10; i++ {
        fmt.Println("Ingrese un número entero: ")
        fmt.Scanln(&numero)
        if numero%5 == 0 {
            contador=contador+1
        }
    }
    fmt.Println("Cantidad de múltiplos de 5: ", contador)
}