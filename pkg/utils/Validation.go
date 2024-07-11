package utils

import (
	"errors"
	"net/url"
	"proxy-server/pkg/entity"
)

func ValidateRequest(req entity.ProxyRequest) error {
	if req.Method == "" {
		return errors.New("method is required")
	}

	if req.URL == "" {
		return errors.New("url is required")
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		return errors.New("invalid url")
	}

	return nil
}
