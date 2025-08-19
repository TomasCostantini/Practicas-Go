/*Guia de ejercicios GO.
Ejercicio 26:
Una empresa de remises ingresa los datos de los viajes realizados por sus móviles. Por cada viaje se
indican el número de móvil (que puede ser del 1 al 6) y la cantidad de cuadras recorridas en el viaje.
El precio por cuadra es de $500.
Se pide:
● Leer los datos de los viajes. El fin de carga viene dado por un número de móvil igual a cero.
● Cargar un vector REM de 6 elementos con el total de cuadras que hizo cada móvil.
● Crear una función que reciba como parámetro un número de móvil y el vector RES e imprima
la cantidad total de cuadras realizadas y el importe recaudado por ese móvil.
● Utilizar la función del punto anterior para imprimir en pantalla el siguiente reporte:
Nro. móvil      Cant. total cuadras            Importe recaudado
1               100                             50000
2               150                             75000
3               90                              45000
4               200                             100000
5               205                             102500
6               90                              45000
● Realizar un procedimiento que permita ordenar un vector de forma descendente. El
procedimiento deberá recibir como parámetro por referencia el vector a ordenar.
● Ordenar el vector REM con el procedimiento desarrollado en el punto anterior e imprimir el
vector ordenado.
*/

package main
import "fmt"
