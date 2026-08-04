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
const(
	indice=14
)
	var(
patente, estado int
)

for i:=0; i<indice;i++{
	fmt.Println("Ingrese el numero de patente: ")
	fmt.Scan(&patente)
	fmt.Println("Ingrese el estado del vehiculo (0 auto funcionando; 1 está en el taller mecánico; 2 está en chapa y pintura y 3 fuera de funcionamiento.)")
	fmt.Scan(&estado)
}
}
