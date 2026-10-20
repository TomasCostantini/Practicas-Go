package main

import "fmt"

const (
	empleados = 15
	columnas=5
)

var ()

func main() {
	var (
		EMP                                                               [empleados]string
		DAT [empleados][columnas] int
		nombre                                                            string
		n_cajero, nro_caja, cantidad_productos, importe, horas_trabajadas int
	)

	for i := 1; i < empleados; i++ {
		fmt.Println("Ingrese el nombre del empleado: ")
		fmt.Scan(&nombre)
		EMP[i] = nombre
	}
	fmt.Println("Ingrese el numero de legajo: ")
	fmt.Scan(&n_cajero)
	for n_cajero != 0 {

		fmt.Println("Ingrese el nro de caja: ")
		fmt.Scan(&nro_caja)
	}
}
