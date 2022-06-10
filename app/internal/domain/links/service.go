package links

import (
	"context"
	"fmt"

	"ozon-url-shortener/app/internal/adapters/links"
	"ozon-url-shortener/app/pkg/generator"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) links.Service {
	return &service{storage: storage}
}

const length = 10

func (s *service) CreateLink(url string) (string, error) {
	key, err := generator.Key(length)
	if err != nil {
		return "", fmt.Errorf("generate key: %w", err)
	}

	err = s.storage.StoreURL(context.Background(), key, url)
	if err != nil {
		return "", fmt.Errorf("store url: %w", err)
	}

	return string(key), nil
}

func (s *service) GetLink(key string) (string, error) {
	url, err := s.storage.GetURL(context.Background(), key)
	if err != nil {
		return "", fmt.Errorf("get url: %w", err)
	}

	return url, nil
}
