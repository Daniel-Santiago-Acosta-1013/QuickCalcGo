package main

import (
    "bufio"
    "fmt"
    "os"
    "QuickCalcGo/pkg/calculator"
    "QuickCalcGo/pkg/parser"
    "github.com/spf13/cobra"
    "strings"
)

var rootCmd = &cobra.Command{
    Use:   "QuickCalcGo",
    Short: "QuickCalcGo es una calculadora CLI avanzada",
    Run: func(cmd *cobra.Command, args []string) {
        // Si no hay argumentos, pasa a la interfaz interactiva
        if len(args) == 0 {
            interactiveMode()
        } else {
            // Calcula la expresión pasada como argumento
            expression := strings.Join(args, " ")
            calculateAndPrint(expression)
        }
    },
}

func interactiveMode() {
    fmt.Println("Modo interactivo. Escribe 'exit' para salir.")
    fmt.Println("Escribe una operación matemática (ej. '2*9' o '3*9/8+9') y presiona Enter.")
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        input := scanner.Text()
        if input == "exit" {
            fmt.Println("Saliendo de QuickCalcGo.")
            break
        }
        calculateAndPrint(input)
    }
}

func calculateAndPrint(input string) {
    expression, err := parser.Parse(input)
    if err != nil {
        fmt.Printf("Error al analizar la expresión: %s\n", err)
        return
    }

    result := calculator.Calculate(expression)
    fmt.Printf("Resultado: %v\n", result)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
