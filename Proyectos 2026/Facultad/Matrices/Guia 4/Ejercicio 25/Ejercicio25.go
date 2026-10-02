/* Guia de ejercicios GO:
Ejercicio 25:
Un comercio necesita obtener información de los cuatro productos que comercializa (numerados del 1 al 4) en
sus cinco zonas de venta (numeradas del 1 al 5). Los datos se ingresan a partir de órdenes de compra que se
van generando de acuerdo al siguiente formato:
Nro de Orden    Zona     Producto1      Producto2     Producto3      Producto4
01               5          22             31            21             18

Donde cada dato representa una venta y se informa la zona donde se realizó y la cantidad vendida de cada uno de
los 4 productos. No se sabe la cantidad de órdenes que se van a ingresar, pero el fin de datos es dado con un cero
en nro. de orden. Se desea obtener:

 El total de ventas por zona y por producto

 Un ranking con los 4 productos ordenados de mayor a menor según la cantidad total de unidades vendidas.
*/

package main

import "fmt"

const (
	columnas = 7
)

func main() {
	var (
		A                                  [][columnas]int
		vector_auxiliar                    [7]int
		numero_orden, zona, p1, p2, p3, p4 int
		suma                               = 0
	)

	for numero_orden > 0 {
		fmt.Println("Ingrese el numero de orden: ")
		fmt.Scan(&numero_orden)
		fmt.Println("Ingrese la zona (1 al 5): ")
		fmt.Scan(&zona)
		fmt.Println("Ingrese la cantidad del producto 1: ")
		fmt.Scan(&p1)
		fmt.Println("Ingrese la cantidad del producto 2: ")
		fmt.Scan(&p2)
		fmt.Println("Ingrese la cantidad del producto 3: ")
		fmt.Scan(&p3)
		fmt.Println("Ingrese la cantidad del producto 4: ")
		fmt.Scan(&p4)
		vector_auxiliar[0] = numero_orden
		vector_auxiliar[1] = zona
		vector_auxiliar[2] = p1
		vector_auxiliar[3] = p2
		vector_auxiliar[4] = p3
		vector_auxiliar[5] = p4
		suma = suma + p1 + p2 + p3 + p4
		vector_auxiliar[6] = suma

		A = append(A, vector_auxiliar)
	}
}
