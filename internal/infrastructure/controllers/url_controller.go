package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/app"
)

type URLOutputDTO struct {
	ShortCode   string `json:"short_code"`
	OriginalURL string `json:"original_url"`
}

type SaveURLInputDTO struct {
	URL string `json:"url"`
}

type URLController struct {
	urlService *app.URLService
}

func NewURLController(urlService *app.URLService) *URLController {
	return &URLController{
		urlService: urlService,
	}
}

func (controller *URLController) Save(response http.ResponseWriter, request *http.Request) {
	var dto SaveURLInputDTO
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if err = controller.urlService.Save(uuid.NewString(), dto.URL); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	response.WriteHeader(http.StatusCreated)
}

func (controller *URLController) List(response http.ResponseWriter, _ *http.Request) {
	urls, err := controller.urlService.List()
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	var dtos []URLOutputDTO
	for _, url := range urls {
		dtos = append(dtos, URLOutputDTO{ShortCode: url.ShortCode(), OriginalURL: url.OriginalURL()})
	}
	encoder := json.NewEncoder(response)
	encoder.Encode(dtos)
}

func (controller *URLController) Redirect(response http.ResponseWriter, request *http.Request) {
	shortCode := request.PathValue("short_code")
	url, err := controller.urlService.Find(shortCode)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(response, request, url.OriginalURL(), http.StatusMovedPermanently)
}
