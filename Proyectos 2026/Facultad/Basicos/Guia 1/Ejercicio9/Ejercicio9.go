/*
Guia de ejercicios GO:
Ejercicio 9:
Leer la cantidad (entera) y precio unitario (real) de un producto, calcular el subtotal (precio neto), IVA
(21%) y total de la factura. Imprimir los resultados.
*/

package main

import "fmt"

var (
	cantidad                              int
	precio_unitario, subtotal, iva, total float32
)

const (
	porcentaje_iva = 21
)

func main() {
	fmt.Println("Ingrese la cantidad de productos: ")
	fmt.Scan(&cantidad)
	fmt.Println("Ingrese el precio unitario: ")
	fmt.Scan(&precio_unitario)

	subtotal = precio_unitario * float32(cantidad)
	total = subtotal + subtotal*porcentaje_iva/100

	fmt.Println("El subtotal de la compra es: ")
	fmt.Println(subtotal)
	fmt.Println("El precio total de la compra es: ")
	fmt.Println(total)
}
