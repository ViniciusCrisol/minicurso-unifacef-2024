package app

import "github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/entities"

type URLDAO interface{}

type URLService struct {
	urlDAO URLDAO
}

func NewURLService(urlDAO URLDAO) URLService {
	return URLService{
		urlDAO: urlDAO,
	}
}

func (service *URLService) Save(shortCode, originalURL string) error {
	return nil
}

func (service *URLService) List() ([]entities.URL, error) {
	return nil, nil
}

func (service *URLService) Find(shortCode string) (entities.URL, error) {
	return entities.URL{}, nil
}
