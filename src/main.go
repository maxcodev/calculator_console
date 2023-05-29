package main

import (
	"bufio"
	"fmt"
	mod "inputConsole/src/models"
	"os"
)

func main() {
	var ops mod.Calc
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Introduce el operador: \na) Suma \nb) Resta \nc) Multiplicación \nd) División")
	scanner.Scan()
	ops.Operator = scanner.Text()

	fmt.Println("Introduce el primer valor:")
	scanner.Scan()
	ops.Val1 = scanner.Text()

	fmt.Println("Introduce el segundo valor valor:")
	scanner.Scan()
	ops.Val2 = scanner.Text()

	value := ops.Operate(ops.Operator, ops.Val1, ops.Val2)

	fmt.Println("Resultado:", value)
}
