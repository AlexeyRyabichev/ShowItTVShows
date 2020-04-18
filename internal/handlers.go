package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func (rt *Router) GetTVShow(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("X-Login")
	showID := r.Header.Get("X-TVShowId")

	show := GetWatchlist(login)[showID]

	resp := TVShow{
		TVShowID: showID,
		Seen:     show.Seen,
		Unseen:   show.Unseen,
		Seasons:  show.Seasons,
	}

	js, err := json.Marshal(&resp)
	if err != nil {
		log.Printf("ERR\tcannot parse TV show to json, %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if bytes, err := w.Write(js); err != nil {
		log.Printf("ERR\t%v", err)
		http.Error(w, "ERR\tcannot write json to response: "+err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Printf("RESP\tGET TVSHOW\twritten %d bytes in response", bytes)
	}
}

func (rt *Router) PostTVShow(w http.ResponseWriter, r *http.Request) {
	type JsonReq struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		TVShowId string `json:"tv_show_id"`
		Seen     bool   `json:"seen"`
		Unseen   bool   `json:"unseen"`
	}
	var jsonReq JsonReq

	if err := json.NewDecoder(r.Body).Decode(&jsonReq); err != nil {
		log.Printf("ERR\t%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	watchlist := GetWatchlist(jsonReq.Login)
	if watchlist[jsonReq.TVShowId] == nil {
		watchlist[jsonReq.TVShowId] = &TVShow{
			TVShowID: jsonReq.TVShowId,
			Seen:     false,
			Unseen:   false,
			Seasons:  nil,
		}
	}

	if jsonReq.Seen {
		watchlist[jsonReq.TVShowId].Seen = true
	}
	if jsonReq.Unseen {
		watchlist[jsonReq.TVShowId].Unseen = true
	}

	if UpdateWatchlist(jsonReq.Login, &watchlist) != true {
		log.Printf("RESP\tPOST SHOW\tcannot update watchlist in db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
