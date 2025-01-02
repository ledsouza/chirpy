package main

import (
	"errors"
	"strings"
)

func CleanBadWords(body string) string {
	badWords :=
		map[string]struct{}{
			"kerfuffle": {},
			"sharbert":  {},
			"fornax":    {},
		}

	cleanedBody := strings.Fields(body)
	for i, word := range cleanedBody {
		if _, ok := badWords[strings.ToLower(word)]; ok {
			cleanedBody[i] = "****"
		}
	}

	return strings.Join(cleanedBody, " ")
}

func ValidateChirp(body string) (string, error) {
	if len(body) > 140 {
		return "", errors.New("Chirp is too long")
	}

	return CleanBadWords(body), nil
}
