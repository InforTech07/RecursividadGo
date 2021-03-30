package main

//import librerias utilizadas en la apliacacion
import (
	"bufio"
	"fmt"
	f "fmt"
	"os"
	"strconv"
	"strings"
)

// var variables globales de la apliacaciòn
//var que almacena el numero para la serie de fibonacci
var numfibo int

//var que almacena el array de vectores que se multiplicaran
var t2 []int

//main funcion principal para el arranque de la aplicacion CLI
func main() {

	app()
}

//app funcion que maneja la recursividad del llamado de las diferentes opciones
func app() {
	var op int
	menu()
	f.Scanln(&op)
	switch op {
	//Opcion que emplea recursividad para en contrar la serie de fibonacci
	case 1:
		titlesection("FIBONACCI", "Ingrese un numero")
		f.Scanln(&numfibo)
		fibonacci(1, 1)
		app()
		//Opcion que emplea la division de numeros por resta
	case 2:
		var dividendo, divisor int
		titlesection("DIVISION", "Ingrese el dividendo")
		f.Scanln(&dividendo)
		f.Println("Ingrese el divisor")
		f.Scanln(&divisor)
		result := divideRecursion(dividendo, divisor)
		f.Println("El resultado es: ", result)
		app()
		//opcion que inverte un numero ingresado
	case 3:
		var num int
		titlesection("INVERTIR-NUMERO", "Ingrese un numero")
		f.Scanln(&num)
		f.Println("Resultado :")
		invertir(num)
		app()
		//opcion que multiplica los elementos de un vector
	case 4:

		titlesection("VECTOR", "Ingrese los valores separados por [,]")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		datainput := scanner.Text()
		r := strings.Split(datainput, ",")
		convertarrayinput(r)
		multarray(t2, len(t2), 0, 1)
		app()
		//opcion que identifica un numero positivo o negativo
	case 5:
		var num float32
		titlesection("NUMERO + o -", "Ingrese un numero")
		f.Scanln(&num)
		positive(num)
		app()
		//opcion para salir de la aplicacion
	case 6:
		os.Exit(0)
		//opcion default
	default:
		f.Println("Opcion no valida")
	}

}

//funciones para cada opcion

// fibonacci que crea la serie de fibonacci apartir de un numero
func fibonacci(m, n int) {
	if m >= numfibo {
		return
	}
	f.Printf("Serie fibonacci: %d\n", m)
	fibonacci(n, (n + m))
}

// divideRecusrion divide dos numeros apartir de resta
func divideRecursion(x, y int) int {
	if y > x {
		return 0
	} else {
		f.Println("---recursion---")
		return 1 + divideRecursion(x-y, y)
	}
}

//invertir inverte un numero ingresado
func invertir(num int) {
	if num >= 10 {
		f.Printf("%d", num%10)
		invertir(num / 10)
	} else {
		f.Printf("%d", num)
		f.Println("")
	}

}

//converarrayinput convierte el strin ingresado en un array de enteros empleando recursividad
func convertarrayinput(in []string) {
	var e []string
	if len(in) == 0 {
		return

	}
	j, err := strconv.Atoi(in[0])
	if err != nil {
		panic(err)
	}
	t2 = append(t2, j)
	e = in[1:]
	convertarrayinput(e)

}

//multarray obtine el producto de los elementos de un array que recibe
func multarray(ve []int, n, f, t int) {

	var tot int = t
	if n == 0 {
		return
	}
	tot = tot * ve[f]
	fmt.Println("producto:", tot)
	multarray(ve, n-1, f+1, tot)
}

//positive identifica si un numero es positivo o neutro
func positive(num float32) {
	if num > 0 {
		f.Println("El numero: ", num, " es positivo")
	} else if num == 0 {
		f.Println("El numero: ", num, " es neutro")
	} else {
		negative(num)
	}
}

//negative identifica si el numero ingresado es negativo
func negative(num float32) {
	if num < 0 {
		f.Println("El numero: ", num, "es negativo")
	} else {
		positive(num)
	}
}

//funciones del cuerpo del CLI
//menu dibuja un menu en la consola.
func menu() {
	t2 = nil
	numfibo = 0
	f.Println("_________________________________________________________________________________")
	f.Println("--------------------------------RECURSIVIDAD-------------------------------------")
	f.Println("[1]-fibonacci [2]-division [3]-invertir [4]-vector [5]-positivonegativo [6]-salir")
	f.Println("_________________________________________________________________________________")
	f.Printf("--> ")
}

//titlesection dibuja un meno por cada opcion ingresado.
func titlesection(ti, i string) {
	f.Println("-------------------[", ti, "]---------------------")
	f.Println("Instrucciones: ", i)
	f.Printf("-> ")
}
