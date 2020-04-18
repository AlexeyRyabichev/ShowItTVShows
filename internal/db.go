package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var username = os.Getenv("DBUSER")
var password = os.Getenv("DBPASS")

func GetWatchlist(login string) Watchlist {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	var watchlistDB WatchlistDB
	err = db.QueryRow(fmt.Sprintf(`select watchlist from tvshows where login='%s'`, login)).Scan(&watchlistDB.Watchlist)
	if err != nil {
		log.Printf("ERR\tcannot get watchlist from db: %v", err)
		return nil
	}

	return DB2Watchlist(&watchlistDB)
}
