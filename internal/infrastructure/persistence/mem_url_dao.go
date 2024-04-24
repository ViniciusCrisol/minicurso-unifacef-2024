package persistence

import "github.com/ViniciusCrisol/minicurso-unifacef-2024/internal/entities"

type MemURLDAO struct {
	urls []entities.URL
}

func NewMemURLDAO() MemURLDAO {
	return MemURLDAO{}
}

func (dao *MemURLDAO) List() ([]entities.URL, error) {
	return dao.urls, nil
}

func (dao *MemURLDAO) Save(url entities.URL) error {
	dao.urls = append(dao.urls, url)
	return nil
}

func (dao *MemURLDAO) FindByShortCode(shortCode string) (entities.URL, bool, error) {
	for _, url := range dao.urls {
		if url.ShortCode() == shortCode {
			return url, true, nil
		}
	}
	return entities.URL{}, false, nil
}

func (dao *MemURLDAO) FindByOriginalURL(originalURL string) (entities.URL, bool, error) {
	for _, url := range dao.urls {
		if url.OriginalURL() == originalURL {
			return url, true, nil
		}
	}
	return entities.URL{}, false, nil
}
