package parser

import (
    "errors"
    "strings"
)

// Parse convierte una cadena de entrada en una estructura de expresión matemática.
func Parse(input string) (string, error) {
    input = strings.TrimSpace(input)
    if input == "" {
        return "", errors.New("la entrada no puede estar vacía")
    }

    // Reemplaza los operadores matemáticos por espacios para el análisis simplificado.
    input = strings.Replace(input, "*", " * ", -1)
    input = strings.Replace(input, "/", " / ", -1)
    input = strings.Replace(input, "+", " + ", -1)
    input = strings.Replace(input, "-", " - ", -1)

    return input, nil
}
