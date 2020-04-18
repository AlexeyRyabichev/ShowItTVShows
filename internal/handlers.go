package internal

import (
	"net/http"
)

func (rt *Router) GetTVShow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) PostTVShow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) DeleteTVShow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) PostSeason(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) DeleteSeason(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) PostSeries(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) DeleteSeries(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rt *Router) GetWatchlist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
