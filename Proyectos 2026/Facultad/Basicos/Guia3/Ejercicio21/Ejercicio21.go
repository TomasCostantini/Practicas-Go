/*Guia de ejercicios GO:
Ejercicio 21:
Se ingresan N pares de números que indica el Número de móvil y la cantidad de cuadras recorridas de una
empresa de Remises. El precio por cuadra es de $1000 y la cantidad de autos son 6, existiendo una
correspondencia con el número de móvil. Se pide:
• Cantidad de viajes realizados por el móvil 3.
• Importe total recaudado por el móvil 5.
• Imprimir el número de móvil que realizó el viaje más largo y el importe recaudado.
*/

package main

import "fmt"

const (
	autos           = 6
	precioporcuadra = 1000
)

var (
	movil, cuadras, movilviajemaslargo int
	cantidadtres                       = 0
	importecinco                       = 0
	i                                  = 0
)

func main() {
	fmt.Println("Ingrese el numero del movil: ")
	fmt.Scan(&movil)
	for movil != 0 {
		fmt.Println("Ingrese las cuadras recorridas: ")
		fmt.Scan(&cuadras)

		if i == 0 {
			movilviajemaslargo = movil
		} else {
			if movil == 5 {
				importecinco = importecinco + (cuadras * precioporcuadra)
			}
			if movil == 3 {
				cantidadtres++
			}
		}
		i++
		fmt.Println("Ingrese el numero del movil: ")
		fmt.Scan(&movil)
	}
	fmt.Println("El movil con el viaje mas largo fue: ", movilviajemaslargo)
	fmt.Println("La cantidad de viajes realizados por el movil 3 es: ", cantidadtres)
	fmt.Println("El importe recaudado por el movil 5 fue de: ", importecinco)
	fmt.Println("Programa finalizado")
}
