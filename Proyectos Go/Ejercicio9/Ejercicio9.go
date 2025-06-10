/* Guia de ejercicios GO
Ejercicio 9: 
Leer la cantidad (entera) y precio unitario (real) de un producto, calcular el subtotal (precio neto), IVA
(21%) y total de la factura. Imprimir los resultados.
*/ 
package main
import "fmt"

func main() {
	var cantidad int
	var precioUnitario float64
	var subtotal, iva, total float64

	// Leer la cantidad y el precio unitario
	fmt.Print("Ingrese la cantidad del producto: ")
	fmt.Scan(&cantidad)
	fmt.Print("Ingrese el precio unitario del producto: ")
	fmt.Scan(&precioUnitario)

	// Calcular el subtotal, IVA y total
	subtotal = float64(cantidad) * precioUnitario
	iva = subtotal * 0.21
	total = subtotal + iva

	// Imprimir los resultados
	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("IVA (21%%): %.2f\n", iva)
	fmt.Printf("Total: %.2f\n", total)
}