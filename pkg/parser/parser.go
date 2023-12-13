package parser

import (
    "errors"
)

// Parse convierte una cadena de entrada en una estructura de expresión matemática.
func Parse(input string) (string, error) {
    if input == "" {
        return "", errors.New("la entrada no puede estar vacía")
    }

    // Aquí se implementaría el análisis de la cadena de entrada.
    // Por ahora, simplemente devolvemos la entrada como está.
    return input, nil
}
