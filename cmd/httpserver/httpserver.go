package main

import (
	"fmt"

	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/helpers"
)

// Para iniciar o programa, execute na pasta raiz do projeto o seguinte comando: go run cmd/httpserver/httpserver.go
func main() {
	randomString := helpers.RandomString()
	fmt.Println(randomString)
}
