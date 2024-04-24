package entities

type URL struct {
	shortCode   string
	originalURL string
}

func NewURL(shortCode, originalURL string) URL {
	return URL{
		shortCode:   shortCode,
		originalURL: originalURL,
	}
}

func (url *URL) ShortCode() string {
	return url.shortCode
}

func (url *URL) OriginalURL() string {
	return url.originalURL
}
