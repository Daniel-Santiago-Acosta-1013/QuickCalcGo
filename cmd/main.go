package main

import (
    "bufio"
    "fmt"
    "os"
    "QuickCalcGo/pkg/calculator"
    "QuickCalcGo/pkg/parser"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("QuickCalcGo - Calculadora CLI")
    fmt.Println("Escribe una operación matemática (ej. '2*9' o '3*9/8+9') y presiona Enter.")

    for {
        fmt.Print("> ")
        scanner.Scan()
        input := scanner.Text()

        if input == "exit" {
            break
        }

        expression, err := parser.Parse(input)
        if err != nil {
            fmt.Printf("Error al analizar la expresión: %s\n", err)
            continue
        }

        result := calculator.Calculate(expression)
        fmt.Printf("Resultado: %v\n", result)
    }
}
