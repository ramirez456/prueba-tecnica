package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return  router
}

var routes = Routes{
	Route {
		"BeerList",
		"GET",
		"/beers",
		BeerList,
	},
	Route {
		"BeerAdd",
		"POST",
		"/beers",
		BeerAdd,
	},
	Route {
		"BeerShow",
		"GET",
		"/beers/{beerID}",
		BeerShow,
	},
	Route {
		"BeerPriceByBox",
		"GET",
		"/beers/{beerID}/boxprice",
		BeerPriceByBox,
	},
}