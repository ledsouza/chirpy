package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (c *apiConfig) handlerCountHit(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	response := fmt.Sprintf(`
	<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>`, c.fileserverHits.Load())
	w.Write([]byte(response))
}

func (c *apiConfig) handlerResetCountHit(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	c.fileserverHits.Store(0)
	response := "Hits reset"
	w.Write([]byte(response))
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Chirp struct {
		Body string `json:"body"`
	}

	type CleanChirp struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	chirp := Chirp{}
	err := decoder.Decode(&chirp)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode JSON", err)
		return
	}

	if len(chirp.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedBody := CleanBadWords(chirp.Body)

	respondWithJSON(w, http.StatusOK, CleanChirp{CleanedBody: cleanedBody})
}
