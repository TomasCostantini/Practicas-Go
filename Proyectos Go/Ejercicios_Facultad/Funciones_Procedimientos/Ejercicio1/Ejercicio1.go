/*
Guia de ejercicios GO
Ejercicio 20:
Leer los siguientes datos: cantidad de dólares a comprar y cantidad de pesos argentinos disponibles y
cotización del día.
Crear una función que devuelva un resultado booleado que indique si la cantidad de pesos con los que
se cuentan alcanzan para efectuar la compra de dólares.
La función recibirá 3 parámetros:
 Cantidad pesos disponibles
 Cantidad de dólar a comprar
 Cotización del día
El fin de ingreso de datos viene dado por la cantidad de pesos = 0.
La cotización del día se debe ingresar una sola vez.
Imprimir por cada juego de datos si alcanza o no para realizar la operación de cambio.
*/
package main

import "fmt"

func main() {
	var (
		cantidad_pesos, cantidad_dolares, cotizacion float64
	)
	fmt.Println("Ingrese la cotizacion del dolar a dia de hoy: ")
	fmt.Scan(&cotizacion)
	fmt.Print("Ingrese la cantidad de pesos que desea cambiar: ")
	fmt.Scan(&cantidad_pesos)
	for cantidad_pesos != 0 {
		fmt.Println("Ingrese la cantidad de dolares que desea comprar: ")
		fmt.Scan(&cantidad_dolares)

		if operacion(cantidad_pesos, cantidad_dolares, cotizacion) == true {
			fmt.Println("La transaccion es valida.")
		} else {
			fmt.Println("La transaccion es invalida.")
		}

		fmt.Println("Ingrese la cantidad de pesos que desea cambiar: ")
		fmt.Scan(&cantidad_pesos)
	}
	fmt.Print("Programa finalizado")
}

func operacion(cantidad_pesos, cantidad_dolares, cotizacion float64) (validez bool) {
	var operacion float64
	operacion = cantidad_dolares * cotizacion
	if operacion <= cantidad_pesos {
		validez = true
	} else {
		validez = false
	}
	return validez
}
