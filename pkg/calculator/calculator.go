package calculator

import (
    "fmt"
    "strconv"
    "strings"
)

// Calculate evalúa una expresión matemática y devuelve el resultado.
func Calculate(expr string) float64 {
    // Aquí iría la lógica para evaluar la expresión.
    // Por simplicidad, vamos a dividir la expresión en partes y realizar una operación básica.
    parts := strings.Fields(expr)
    if len(parts) != 3 {
        fmt.Println("Formato de expresión no soportado.")
        return 0
    }

    num1, _ := strconv.ParseFloat(parts[0], 64)
    operator := parts[1]
    num2, _ := strconv.ParseFloat(parts[2], 64)

    switch operator {
    case "+":
        return num1 + num2
    case "-":
        return num1 - num2
    case "*":
        return num1 * num2
    case "/":
        return num1 / num2
    default:
        fmt.Println("Operador desconocido:", operator)
        return 0
    }
}
