package calculator

import (
    "github.com/Knetic/govaluate"
    "log"
)

// Calculate evalúa una expresión matemática y devuelve el resultado.
func Calculate(expr string) float64 {
    expression, err := govaluate.NewEvaluableExpression(expr)
    if err != nil {
        log.Fatalf("Error al crear la expresión: %s", err)
    }

    result, err := expression.Evaluate(nil)
    if err != nil {
        log.Fatalf("Error al evaluar la expresión: %s", err)
    }

    return result.(float64)
}
