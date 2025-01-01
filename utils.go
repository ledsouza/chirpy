package main

import "strings"

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
