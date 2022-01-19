package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)
var beers = Beers{
	Beer{
		1,
		"Cristal",
		12.52,
		"Cristal",
	},
	Beer{
		2,
		"Pilse",
		12.52,
		"Pilse",
	},
	Beer{
		3,
		"Cuzqueña",
		12.52,
		"Cuzqueña",
	},
	Beer{
		4,
		"Budweiser",
		12.52,
		"Budweiser",
	},
	Beer{
		5,
		"Corona",
		12.52,
		"Corona",
	},
}

var collection = getSession().DB("prueba").C("beers")
func getSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost")

	if(err != nil){
		panic(err)
	}

	return session
}

func BeerList(w http.ResponseWriter, r *http.Request){

	json.NewEncoder(w).Encode(beers)
}

func BeerAdd(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)

	var beer_data Beer

	err := decoder.Decode(&beer_data)

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close()

	log.Println(beer_data)
	beers =  append(beers, beer_data)

	collection.Insert(beer_data)
}

func BeerShow(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Un chela específica")
}

func BeerPriceByBox(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Precio por Caja")
}

