package persistence

import "github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/entities"

type MemURLDAO struct{}

func NewMemURLDAO() MemURLDAO {
	return MemURLDAO{}
}

func (dao *MemURLDAO) List() ([]entities.URL, error) {
	return nil, nil
}

func (dao *MemURLDAO) Save(url entities.URL) error {
	return nil
}

func (dao *MemURLDAO) FindByShortCode(shortCode string) (entities.URL, bool, error) {
	return entities.URL{}, false, nil
}

func (dao *MemURLDAO) FindByOriginalURL(originalURL string) (entities.URL, bool, error) {
	return entities.URL{}, false, nil
}
