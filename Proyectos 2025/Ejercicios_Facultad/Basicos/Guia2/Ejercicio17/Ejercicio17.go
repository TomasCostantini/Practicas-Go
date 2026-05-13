/*
En una elección existen 3 candidatos. Por cada elector se tienen los siguientes datos: sexo (H o M) y el
partido por el cual votó (1 a 3). Ingresar “n” electores, determinar e imprimir:
• La cantidad de mujeres que votaron.
• La cantidad de varones que votaron al candidato 3.
• Candidato que ganó la elección, considerando que no hay empate.
IMPORTANTE: El valor del sexo se debe ingresar en mayusculas.
*/
package main

import "fmt"

func main() {

	var (
		sexo      string
		partido   int
		mujeres   int
		cantidad3 int
		n         int
		votos1    int
		votos2    int
		votos3    int
		ganador   int
	)

	fmt.Print("Ingrese la cantidad de electores: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Ingrese el sexo (H o M): ")
		fmt.Scan(&sexo)
		fmt.Print("Ingrese el partido (1 a 3): ")
		fmt.Scan(&partido)

		if partido == 1 {
			votos1++
		} else if partido == 2 {
			votos2++
		} else if partido == 3 {
			votos3++
		}

		if sexo == "M" {
			mujeres++
		} else if sexo == "H" && partido == 3 {
			cantidad3++
		}

		if votos1 > votos2 && votos1 > votos3 {
			ganador = 1
		} else if votos2 > votos1 && votos2 > votos3 {
			ganador = 2
		} else if votos3 > votos1 && votos3 > votos2 {
			ganador = 3
		}
	}

	fmt.Println("Cantidad de mujeres que votaron: ", mujeres)
	fmt.Println("Cantidad de varones que votaron al candidato 3: ", cantidad3)
	fmt.Println("Candidato que ganó la elección:", ganador)
}
