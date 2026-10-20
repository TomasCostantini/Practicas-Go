package main
import "fmt"


const(
	empleados=15
)

var(

)

func main(){
var(
	EMP[empleados] string
	nombre string
)
for i:=1;i<empleados;i++{
fmt.Println("Ingrese el nombre del empleado: ")
fmt.Scan(&nombre)
EMP[i]=nombre
}
}