package app

import (
	"errors"

	"github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/entities"
)

var (
	ErrURLNotFound             = errors.New("url not found")
	ErrShortCodeAlreadyInUse   = errors.New("short code already in use")
	ErrOriginalURLAlreadyInUse = errors.New("original url already in use")
)

type URLDAO interface {
	Save(url entities.URL) error
	List() ([]entities.URL, error)
	FindByShortCode(shortCode string) (entities.URL, bool, error)
	FindByOriginalURL(originalURL string) (entities.URL, bool, error)
}

type URLService struct {
	urlDAO URLDAO
}

func NewURLService(urlDAO URLDAO) URLService {
	return URLService{
		urlDAO: urlDAO,
	}
}

func (service *URLService) Save(shortCode, originalURL string) error {
	_, found, err := service.urlDAO.FindByShortCode(shortCode)
	if err != nil {
		return err
	}
	if found {
		return ErrShortCodeAlreadyInUse
	}
	_, found, err = service.urlDAO.FindByOriginalURL(originalURL)
	if err != nil {
		return err
	}
	if found {
		return ErrOriginalURLAlreadyInUse
	}
	return service.urlDAO.Save(entities.NewURL(shortCode, originalURL))
}

func (service *URLService) List() ([]entities.URL, error) {
	return service.urlDAO.List()
}

func (service *URLService) Find(shortCode string) (entities.URL, error) {
	url, found, err := service.urlDAO.FindByShortCode(shortCode)
	if err != nil {
		return entities.URL{}, err
	}
	if !found {
		return entities.URL{}, ErrURLNotFound
	}
	return url, nil
}
