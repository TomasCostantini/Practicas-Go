/*Guia de ejercicios GO:
Ejercicio 29:
Una compañía. de transporte tiene almacenado en un vector de 14 elementos la siguiente información:
Patente Estado.
El elemento estado puede tomar alguno de los siguientes valores: 0: auto funcionando; 1: está en el taller
mecánico; 2: está en chapa y pintura y 3: fuera de funcionamiento.
Se pide:
a) Cargar el vector con los datos de Patente y Estado
b) Determinar la cantidad de vehículos que tienen estado 0 (cero), 1, 2y 3.
c) Generar un vector con la patente de aquellos vehículos que tengan estado igual a 0
d) Imprimir el vector generado.
*/

package main

import "fmt"

func main() {
	const (
		indice = 14
	)
	var (
		vector          [14]int
		estadocero      []int
		patente, estado int
		cero            = 0
		uno             = 0
		dos             = 0
		tres            = 0
	)

	for i := 0; i < indice-1; i++ {
		fmt.Println("Ingrese el numero de patente: ")
		fmt.Scan(&patente)
		fmt.Println("Ingrese el estado del vehiculo (0 auto funcionando; 1 está en el taller mecánico; 2 está en chapa y pintura y 3 fuera de funcionamiento.)")
		fmt.Scan(&estado)
		vector[i] = patente
		vector[i+1] = estado
	}
	for i := 0; i < indice; i++ {
		switch {
		case vector[i+1] == 0:
			cero = cero + 1
			estadocero = append(estadocero, vector[i])
		case vector[i+1] == 1:
			uno = uno + 1

		case vector[i+1] == 2:
			dos = dos + 1

		case vector[i+1] == 3:
			tres = tres + 1

		}
	}

	fmt.Println("Cantidad de vehiculos en estado 1: ", uno)
	fmt.Println("Cantidad de vehiculos en estado 2: ", dos)
	fmt.Println("Cantidad de vehiculos en estado 3: ", tres)
	fmt.Println("Cantidad de vehiculos en estado 0: ", cero)
	for i := 0; i < indice; i++ {
		fmt.Print(estadocero[i])
		fmt.Print(" , ")
	}
}
