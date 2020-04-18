package internal

import "encoding/json"

//type WatchlistHTTP struct {
//	Login     string    `json:"login"`
//	Watchlist Watchlist `json:"watchlist"`
//}

type WatchlistDB struct {
	Login     string `db:"login"`
	Watchlist string `db:"watchlist"`
}

type WatchlistResponse struct {
	Login         string   `json:"login"`
	SeenTVShows   []string `json:"seen_tv_shows"`
	UnseenTVShows []string `json:"unseen_tv_shows"`
}

type Watchlist map[string]*TVShow

type TVShow struct {
	TVShowID string    `json:"tv_show_id"`
	Seen     bool      `json:"seen"`
	Unseen   bool      `json:"unseen"`
	Seasons  []Seasons `json:"seasons"`
}

type Seasons struct {
	Season int      `json:"season"`
	Seen   bool     `json:"seen"`
	Series []Series `json:"series"`
}

type Series struct {
	SeriesID string `json:"series_id"`
	Seen     bool   `json:"seen"`
}

func DB2Watchlist(db *WatchlistDB) Watchlist {
	var watchlist Watchlist
	if err := json.Unmarshal([]byte(db.Watchlist), &watchlist); err != nil {
		return nil
	}

	//watchlist := make(Watchlist)
	//for _, tvshow := range watchlistArr {
	//	watchlist[tvshow.TVShowID] = &tvshow
	//}
	return watchlist
}

func Watchlist2DB(watchlist *Watchlist) (WatchlistDB, error) {
	js, err := json.Marshal(watchlist)
	if err != nil {
		return WatchlistDB{}, err
	}

	return WatchlistDB{Watchlist: string(js)}, nil
}
