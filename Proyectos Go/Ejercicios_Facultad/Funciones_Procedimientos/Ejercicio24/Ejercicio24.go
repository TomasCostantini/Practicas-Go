/* Guia de ejercicios GO.
Ejercicio 24:
Realizar un programa que permita jugar un número indeterminado de veces a “Adivina el Número”
que consiste en lo siguiente:
• Genere un número aleatoriamente comprendido entre 0-100.
• Permitir al jugador adivinar el número generado, para ello deberá introducir por teclado números
mientras no acierte. Para ello desarrollar la función “checkNumeroAdivinado” que valide si el
usuario adivinó el número o no. La función deberá recibir como parámetro el número ingresado
por el usuario y el número generado aleatoriamente y devolver como resultado:
o -1 si el número ingresado es menor al número adivinar.
o 1 si el número ingresado es mayor al número adivinar.
o 0 si el número ingresado es igual al número adivinar.
• Mientras el número no sea adivinado, se deberá seguir leyendo un nuevo número hasta que sea
adivinado.
• Contar el número de intentos realizados.
• Realizar un procedimiento que según el resultado de la función “checkNumeroAdivinado”,
imprima en pantalla si el número generado es mayor, menor o igual que el ingresado. El
procedimiento deberá recibir como parámentro el resultado de la función
“checkNumeroAdivinado” y la cantidad de intentos realizados.
En el momento de la victoria, además se deberá indicar según el número de intentos realizados:
<= 5 se indica al jugador que es bueno.
5 < número de intentos < 15 se indica al jugador que es regular.
el número de intentos es > 15 se indica que el jugador no es muy bueno.

Generación de números aleatorios.
Para generar números aleatorios, Go cuenta con la función rand.Intn(nroMax) que se encuentra en el
paquete “math/rand”.
La función rand.Intn genera un número aleatorio de tipo entero, comprendido entre 0 el número máximo
pasado como parámetro (nroMax).
Sin embargo el uso de la función rand.Intn es en ocasiones no es insuficiente por si sola ya que los números
son realmente pseudoaleatorios, y cada vez que se ejecute el programa se obtendrían los mismos valores.
Para evitar esto debemos utilizar la función rand.Seed(time.Now().Unix()), que inicializa el generador de
números aleatorios, y asegura que los números que obtendrá serán diferentes cada vez que se ejecute el
programa. Para ello poner en uso el paquete “time”.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	numeroAdivinar := rand.Intn(101)
	var numeroIngresado int
	var intentos int

	for {
		fmt.Print("Ingrese un número entre 0 y 100: ")
		fmt.Scan(&numeroIngresado)
		resultado := checkNumeroAdivinado(numeroIngresado, numeroAdivinar)
		intentos++

		if resultado == 0 {
			fmt.Println("Adivinaste el número.")
			break
		} else {
			imprimirResultado(resultado, intentos)
		}
	}
}

func checkNumeroAdivinado(numeroIngresado, numeroAdivinar int) int {
	if numeroIngresado < numeroAdivinar {
		return -1
	} else if numeroIngresado > numeroAdivinar {
		return 1
	} else {
		return 0
	}
}

func imprimirResultado(resultado, intentos int) {
	switch resultado {
	case -1:
		fmt.Println("El número a adivinar es mayor.")
	case 1:
		fmt.Println("El número a adivinar es menor.")
	}
	if intentos <= 5 {
		fmt.Println("¡Eres un jugador excelente!")
	} else if intentos <= 15 {
		fmt.Println("¡Eres un jugador promedio!")
	} else {
		fmt.Println("¡Necesitas mejorar!")
	}
}
