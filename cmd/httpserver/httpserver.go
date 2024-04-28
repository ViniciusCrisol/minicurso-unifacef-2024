package main

import (
	"net/http"

	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/app"
	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/infrastructure/controllers"
	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/infrastructure/persistence"
)

// Para iniciar o programa, execute na pasta raiz do projeto o seguinte comando: go run cmd/httpserver/httpserver.go
func main() {
	memURLDAO := persistence.NewMemURLDAO()
	urlService := app.NewURLService(&memURLDAO)
	urlController := controllers.NewURLController(&urlService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/urls", urlController.List)
	mux.HandleFunc("POST /api/urls", urlController.Save)
	mux.HandleFunc("GET /{short_code}", urlController.Redirect)
	http.ListenAndServe(":8080", mux)
}
