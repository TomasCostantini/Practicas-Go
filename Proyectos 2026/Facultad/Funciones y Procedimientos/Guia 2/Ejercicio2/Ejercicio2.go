/*
Guia de ejercicios GO:
Ejercicio 2:
Realizar un programa que permita jugar un número indeterminado de veces a “Adivina el Número”
que consiste en lo siguiente:
● Genere un número aleatoriamente comprendido entre 0-100.
● Permitir al jugador adivinar el número generado, para ello deberá introducir por teclado
números mientras no acierte. Desarrollar una función que valide si el usuario adivinó el
número o no, la función deberá recibir como parámetro el número ingresado por el
usuario y el número generado aleatoriamente e informar por pantalla si el número
generado es mayor o menor que el ingresado. Además deberá devolver un valor
booleando que indique si acertó el número o no.

●Contar el número de intentos.
●En el momento de la victoria, mostrarlo en pantalla e informar el número de intentos que ha
utilizado, sí el número de intentos es:

o <= 5 se indica al jugador que es bueno.
o 5 < número de intentos < 15 se indica al jugador que es regular.
o el número de intentos es > 15 se indica que el jugador no es muy bueno.
*/
package main

import "fmt"
import "time"
import "math/rand"

var (
	validez  bool
	intentos = 0
)

func main() {
	const (
		maximo = 120
	)
	var (
		numero, numero_aleatorio int
	)
	rand.Seed(time.Now().Unix())
	numero_aleatorio = rand.Intn(maximo)


	
	for validez != false {
		fmt.Println("Ingrese un numero: ")
		fmt.Scan(&numero)
		verificacion(numero, numero_aleatorio)
	}
	
	fmt.Println("El numero de intentos fue: ", intentos)
}

func verificacion(numero, aleatorio int) bool {
	switch {
	case numero < aleatorio:
		validez = false
		fmt.Println("El numero es mayor")
	case numero == aleatorio:
		validez = true
	case numero > aleatorio:
		validez = false
		fmt.Println("El numero es menor")
	}
	intentos++
	return validez
}
