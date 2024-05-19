package controllers

import (
	"net/http"

	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/app"
)

type URLController struct {
	urlService *app.URLService
}

func NewURLController(urlService *app.URLService) *URLController {
	return &URLController{
		urlService: urlService,
	}
}

func (controller *URLController) Save(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func (controller *URLController) List(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func (controller *URLController) Redirect(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, "https://oglobo.globo.com", http.StatusMovedPermanently)
}
