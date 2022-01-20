package main

import (
	"gopkg.in/mgo.v2/bson"
)

type Beer struct {
	ID  bson.ObjectId `bson:"_id"`
	Name string `json:"name"`
	Brewery string `json: "brewery"`
	Country string `json: "country"`
	Price float64 `json: "price"`
	Currency string `json: "currency"`
}

type Beers []Beer

type PriceBeerBox struct {
	PriceTotal float64
}