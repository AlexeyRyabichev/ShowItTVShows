package internal

import (
	"encoding/json"
	"log"
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
	login := r.Header.Get("X-Login")

	watchlist := GetWatchlist(login)
	//if watchlist == nil {
	//	log.Printf("RESP\tGET WATCHLIST\tcannot find watchlist for user %s", login)
	//	w.WriteHeader(http.StatusBadRequest)
	//}

	watchlistResponse := WatchlistResponse{
		Login:         login,
		SeenTVShows:   []string{},
		UnseenTVShows: []string{},
	}

	for _, tvshow := range watchlist {
		if tvshow.Seen {
			watchlistResponse.SeenTVShows = append(watchlistResponse.SeenTVShows, tvshow.TVShowID)
		}
		if tvshow.Unseen {
			watchlistResponse.UnseenTVShows = append(watchlistResponse.UnseenTVShows, tvshow.TVShowID)
		}
	}

	js, err := json.Marshal(&watchlistResponse)
	if err != nil {
		log.Printf("ERR\tcannot parse watchlist to json, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if bytes, err := w.Write(js); err != nil {
		log.Printf("ERR\t%v", err)
		http.Error(w, "ERR\tcannot write json to response: "+err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Printf("RESP\tGET WATCHLIST\twritten %d bytes in response", bytes)
	}
}
