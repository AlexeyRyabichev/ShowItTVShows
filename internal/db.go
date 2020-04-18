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

func UpdateWatchlist(login string, watchlist *Watchlist) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	dbWatchlist, err := Watchlist2DB(watchlist)
	if err != nil {
		log.Printf("ERR\t%v", err)
		return false
	}

	if _, err := db.Exec("update tvshows set watchlist=$1 where login=$2", dbWatchlist.Watchlist, login); err != nil {
		log.Printf("cannot insert watchlist in table: %v", err)
		return false
	}
	return true
}