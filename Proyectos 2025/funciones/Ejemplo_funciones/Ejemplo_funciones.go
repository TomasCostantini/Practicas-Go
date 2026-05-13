package main
import "fmt"

func main() {
	var (
		n1, n2, resultado int
	)

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&n1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&n2)

	resultado = sumar(n1, n2)
	fmt.Println("La suma de los numeros es: ", resultado)

}

func sumar(a, b int) int {
	suma := a + b
	return suma
}
