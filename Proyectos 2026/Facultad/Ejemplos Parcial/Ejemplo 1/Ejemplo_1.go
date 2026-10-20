package main

import "fmt"

const (
	empleados = 15
	columnas  = 5
)

var ()

func main() {
	var (
		EMP                                                               [empleados]string
		DAT                                                               [empleados][columnas]int
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
		DAT[n_cajero][0] = n_cajero
		fmt.Println("Ingrese el nro de caja: ")
		fmt.Scan(&nro_caja)
		DAT[n_cajero][1] = nro_caja
		fmt.Println("Ingrese la cantidad de productos facturados: ")
		fmt.Scan(&cantidad_productos)
		DAT[n_cajero][2] = cantidad_productos
		fmt.Println("Ingrese el importe facturado: ")
		fmt.Scan(&importe)
		DAT[n_cajero][3] = importe
		fmt.Println("Ingrese la cantidad de horas trabajadas: ")
		fmt.Scan(&horas_trabajadas)
		DAT[n_cajero][4] = horas_trabajadas
		fmt.Println("Ingrese el numero de legajo: ")
		fmt.Scan(&n_cajero)
	}
}
