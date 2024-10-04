package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-rate-limiter/dto"
	"net/http"
)

type UrlController struct {
	db *map[string]string
}

func NewUrlController(db *map[string]string) UrlController {
	return UrlController{db: db}
}

func (c UrlController) CreateUrl(w http.ResponseWriter, r *http.Request)  {
	var url dto.CreateUrlRequestDto
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash := sha256.New()
	hash.Write([]byte(url.Url))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	trucatedHash := hashString[:8]

	(*c.db)[trucatedHash] = url.Url
	
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dto.CreateUrlResponseDto{ShortUrl: fmt.Sprintf("http://localhost:3000/%s", trucatedHash)})
	if err != nil {
		return 
	}
}

func (c UrlController) GetUrl(w http.ResponseWriter, r *http.Request)  {
	urlId := chi.URLParam(r, "urlId")
	url, ok := (*c.db)[urlId]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
