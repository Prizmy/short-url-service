package storage

import (
	"errors"
	"short-url-service/pkg/shortener"
	"sync"
)

type InMemoryStorage struct {
	sUrl            map[string]string
	shortToOriginal map[string]string
	mu              sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		sUrl:            make(map[string]string),
		shortToOriginal: make(map[string]string),
	}
}

func (s *InMemoryStorage) Post(url string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if shortUrl, exists := s.sUrl[url]; exists {
		return shortUrl, nil
	}
	shortURL := shortener.Generate()
	s.sUrl[url] = shortURL
	s.shortToOriginal[shortURL] = url
	return shortURL, nil
}

func (s *InMemoryStorage) Get(shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, exists := s.shortToOriginal[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}
	return url, nil
}
