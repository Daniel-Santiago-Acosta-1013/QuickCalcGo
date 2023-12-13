package calculator

import (
    "fmt"
    "math"
)

// Calculate evalúa una expresión matemática y devuelve el resultado.
func Calculate(expr string) float64 {
    // Aquí iría la lógica para evaluar la expresión.
    // Por simplicidad, vamos a devolver un valor fijo.
    // Implementar la lógica real requeriría un analizador de expresiones más complejo.
    fmt.Println("Evaluando expresión:", expr)
    return math.Sqrt(16) // Devuelve un resultado de ejemplo.
}
