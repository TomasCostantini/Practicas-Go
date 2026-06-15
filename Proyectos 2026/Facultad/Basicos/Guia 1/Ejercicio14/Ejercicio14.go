/*
Guia de ejercicios GO:
Ejercicio 14:
Leer un número entero e imprimir la suma de todos los números anteriores. Por ejemplo, si se ingresa el
número 4, imprimir el resultado de: 1 + 2 + 3.
*/
package maim

import "fmt"

var (
	numero, suma int
)

func main(){
	fmt.Println("Ingrese un numero entero: ")
	fmt.Scan(&numero)

	for i := numero; i >= 0; i-- {
		suma = suma + i
	}
	fmt.Println("La suma total es: ")
	fmt.Println(suma)
	fmt.Println("Programa finalizado")
}
