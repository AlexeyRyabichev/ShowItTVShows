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
	Series   []Episode `json:"series"`
}

type Episode struct {
	SeriesID string `json:"series_id"`
	Seen     bool   `json:"seen"`
}

type TVShowLocal struct {
	TVShowID string `json:"tv_show_id"`
	Seen     bool
	Unseen   bool
	Episodes map[string]*Episode
}

func TVShow2Local(show *TVShow) *TVShowLocal {
	local := TVShowLocal{
		TVShowID: show.TVShowID,
		Seen:     show.Seen,
		Unseen:   show.Unseen,
		Episodes: make(map[string]*Episode),
	}

	for _, episode := range show.Series {
		if local.Episodes[episode.SeriesID] == nil {
			local.Episodes[episode.SeriesID] = &episode
		}
	}

	return &local
}

func Local2TVShow(local TVShowLocal) *TVShow {
	show := TVShow{
		TVShowID: local.TVShowID,
		Seen:     local.Seen,
		Unseen:   local.Unseen,
		Series:   []Episode{},
	}

	for _, episode := range local.Episodes {
		show.Series = append(show.Series, *episode)
	}

	return &show
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
