package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	token, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found {
		return "", errors.New("malformed authorization header")
	}

	token = strings.TrimSpace(token)
	if token == "" {
		return "", errors.New("bearer token missing")
	}

	return token, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	apiKey, found := strings.CutPrefix(authHeader, "ApiKey ")
	if !found {
		return "", errors.New("malformed authorization header")
	}

	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return "", errors.New("api key missing")
	}

	return apiKey, nil
}
