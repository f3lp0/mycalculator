// Nombre de la clase o tema
// se elimina main porque estamos haciendo un módulo, main es lo que
// se ejecuta al inicio, este módulo no se ejecuta al inicio
package mycalculator

import (
	"bufio" // puede crear un scanner para leer inputs del teclado
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

// Receiver function, estamos creando un método a calc
// aplicación: calc.operate(s1,s2)
func (Calc) Pperate(entrada string, operador string) int {
	// Si operador 1 u operador2 no son numéricos,
	// GO toma solo los numéricos
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println("Haciendo la suma numérica:")
		return (operador1 + operador2)
	case "-":
		fmt.Println("Haciendo la resta numérica:")
		return (operador1 - operador2)
	case "*":
		fmt.Println("Haciendo la multiplicación numérica:")
		return (operador1 * operador2)
	case "/":
		fmt.Println("Haciendo la división numérica:")
		return (operador1 / operador2)
	default:
		fmt.Println("Operador no soportado...")
		return 0
	}
}

func parsear(entrada string) int {
	// haciendo la operación con enteros
	// err es una función de error, puede ser nil
	// Empleando wild card sería con _ en luga de err
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func LeerEntrada() string {
	// Creando el input desde el teclado
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// func main() {
// 	entrada := LeerEntrada()
// 	operador := LeerEntrada()

// 	// mostrando los datos de entrada
// 	// fmt.Println(entrada)
// 	// fmt.Println(operador)

// 	c := Calc{}
// 	fmt.Println(c.Operate(entrada, operador))
// }
