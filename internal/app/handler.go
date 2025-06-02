package app

import (
	"io"
	"math/rand"
	"net/http"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var urls = make(map[string]string)
var ID string

func generateID(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func LinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" && r.Method == http.MethodPost {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		}

		bodyText := string(bodyBytes)
		w.Header().Set("Content-type", "text/plain")
		ID = generateID(10)
		urls[ID] = bodyText
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://localhost:8080" + "/" + ID))
	} else if r.Method == http.MethodGet {
		ID = r.URL.Path[1:]
		value, exist := urls[ID]
		if exist {
			http.Redirect(w, r, value, http.StatusTemporaryRedirect)
		} else {
			http.Error(w, "Неккоректный запрос", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Неккоректный запрос", http.StatusBadRequest)
	}
}
