package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"github.com/gorilla/mux"
)

type UrlShortener struct {
	//urls map[string]string
}

func NewUrlShortener() *UrlShortener {
	m := new(UrlShortener)
	//m.urls = make(map[string]string)
	return m
}

func (this UrlShortener) addUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type Message struct {
		Url string `json:"url"`
	}

	var msg Message
	json.NewDecoder(r.Body).Decode(&msg)

	type Answer struct {
		Key string `json:"key"`
	}

	ans := Answer{base64.RawURLEncoding.EncodeToString([]byte(msg.Url))}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(ans)
	fmt.Fprintln(w, string(data))
}

func (this UrlShortener) getUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	Url, _ := base64.RawURLEncoding.DecodeString(key)

	w.Header().Set("Location", string(Url))
	w.WriteHeader(http.StatusMovedPermanently)
}

func main() {
	r := mux.NewRouter()
	service := NewUrlShortener()
	r.HandleFunc("/", service.addUrl)
	r.HandleFunc("/{key}", service.getUrl)
	http.ListenAndServe(":8082", r)
}
