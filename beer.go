package main

type Beer struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Price float64 `json: "price"`
	Brand string `json: "brand"`
}

type Beers []Beer