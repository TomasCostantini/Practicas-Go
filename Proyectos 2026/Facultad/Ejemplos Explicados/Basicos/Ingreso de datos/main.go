package main
import "fmt"


func main(){                 
var(                         //
	numero int               // Seccion de declaracion de variables
	letra, simbolo string    //
)                            //          


fmt.Println("Ingrese un numero: ")    //
fmt.Scan(&numero)                     //
fmt.Println("Ingrese una letra: ")    //
fmt.Scan(&letra)                      //   
fmt.Println("Ingrese un simbolo: ")   //
fmt.Scan(&simbolo)                    //

fmt.Println("El numero es numero: ")  //
fmt.Println(numero)                   //
fmt.Println("La letra es: ")          //
fmt.Println(letra)                    //
fmt.Println("El simbolo es: ")        //
fmt.Println(simbolo)                  //
}
