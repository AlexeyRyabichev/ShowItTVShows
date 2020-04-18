package internal

import (
	"github.com/AlexeyRyabichev/ShowItGate"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Pattern     string           `json:"pattern"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

type Router struct {
	cfg    ShowItGate.NodeCfg
	routes []Route

	Router *mux.Router
}

func NewRouter(cfg ShowItGate.NodeCfg) *Router {
	router := Router{
		cfg: cfg,
	}
	router.routes = []Route{
		{
			Name: "Get TV Show info",
			Method: "GET",
			Pattern: "/v1/tvshow/",
			HandlerFunc: router.GetTVShow,
		},
		{
			Name: "Update TV Show info in watchlist",
			Method: "POST",
			Pattern: "/v1/tvshow/",
			HandlerFunc: router.PostTVShow,
		},
		{
			Name: "Delete TV Show from watchlist",
			Method: "DELETE",
			Pattern: "/v1/tvshow/",
			HandlerFunc: router.DeleteTVShow,
		},
		{
			Name: "Add episode to watchlist",
			Method: "POST",
			Pattern: "/v1/tvshow/series",
			HandlerFunc: router.PostSeries,
		},
		{
			Name: "Delete episode from watchlist",
			Method: "DELETE",
			Pattern: "/v1/tvshow/series",
			HandlerFunc: router.DeleteSeries,
		},
		{
			Name: "Get watchlist",
			Method: "GET",
			Pattern: "/v1/tvshow/watchlist",
			HandlerFunc: router.GetWatchlist,
		},
		{
			Name: "Get TV Show info",
			Method: "GET",
			Pattern: "/v2/tvshow/",
			HandlerFunc: router.GetTVShow,
		},
		{
			Name: "Update TV Show info in watchlist",
			Method: "POST",
			Pattern: "/v2/tvshow/",
			HandlerFunc: router.PostTVShow,
		},
		{
			Name: "Delete TV Show from watchlist",
			Method: "DELETE",
			Pattern: "/v2/tvshow/",
			HandlerFunc: router.DeleteTVShow,
		},
		{
			Name: "Add episode to watchlist",
			Method: "POST",
			Pattern: "/v2/tvshow/series",
			HandlerFunc: router.PostSeries,
		},
		{
			Name: "Delete episode from watchlist",
			Method: "DELETE",
			Pattern: "/v2/tvshow/series",
			HandlerFunc: router.DeleteSeries,
		},
		{
			Name: "Get watchlist",
			Method: "GET",
			Pattern: "/v2/tvshow/watchlist",
			HandlerFunc: router.GetWatchlist,
		},
	}
	router.initRouter()
	return &router
}

func (rt *Router) initRouter() {
	rt.Router = mux.NewRouter().StrictSlash(true)

	for _, route := range rt.routes {
		rt.addRoute(route)
	}

	rt.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"HANDLER NOT FOUND FOR REQUEST: %s %s",
			r.Method,
			r.RequestURI,
		)
		w.WriteHeader(http.StatusNotFound)
	})

	//rt.Router.Use(mux.CORSMethodMiddleware(rt.Router))
}

func (rt *Router) addRoute(route Route) {
	var handler http.Handler
	handler = route.HandlerFunc
	handler = ShowItGate.Logger(handler, route.Name)

	rt.Router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)

	//rt.Router.HandleFunc(route.Pattern, func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	w.Header().Set("Access-Control-Allow-Methods", "*")
	//	w.Header().Set("Access-Control-Allow-Headers", "*")
	//}).Methods(http.MethodOptions)
}
