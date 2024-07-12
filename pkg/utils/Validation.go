package utils

import (
	"errors"
	"net/http"
	"proxy-server/pkg/entity"
)

var allowedMethods = map[string]bool{
	http.MethodGet:    true,
	http.MethodPost:   true,
	http.MethodPut:    true,
	http.MethodDelete: true,
	http.MethodPatch:  true,
}

func ValidateRequest(req entity.ProxyRequest) error {
	if req.Method == "" {
		return errors.New("method is required")
	}
	if !allowedMethods[req.Method] {
		return errors.New("method is not allowed")
	}
	if req.URL == "" {
		return errors.New("url is required")
	}
	return nil
}
