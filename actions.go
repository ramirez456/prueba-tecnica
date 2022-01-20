package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net"
	"net/http"
	"strconv"
)

func createConnection() (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs:    []string{
			"tuchamba-shard-00-00.ikhxh.mongodb.net:27017",
			"tuchamba-shard-00-01.ikhxh.mongodb.net:27017",
			"tuchamba-shard-00-02.ikhxh.mongodb.net:27017",
		},
		Username: "tuchamba",
		Password: "ebcnemj987",
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	return mgo.DialWithInfo(&dialInfo)
}

func responseBeer(w http.ResponseWriter, status int, results Beer) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func index(w http.ResponseWriter, r *http.Request) {


	data := []string{
			"Name: Max Houston Ramirez Martel",
			"DNI: 46972239",
			"Listar:  GET http://localhost:8000/beers",
			"Crear POST http://localhost:8000/beers",
			"Buscar por ID GET http://localhost:8000/beers/61e99038d15da26053b135c1",
			"Precio Caja http://localhost:8000/beers/61e99038d15da26053b135c1/boxprice/0/pen",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)


}
func searchBeers(w http.ResponseWriter, r *http.Request){

	session := mongoConn.Copy()
	defer session.Close()

	entities := []Beer{}
	err := session.DB("prueba").C("beers").Find(bson.M{}).Sort("-_id").All(&entities)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(entities)
}

func addBeers(w http.ResponseWriter, r *http.Request) {
	session := mongoConn.Copy()
	defer session.Close()

	decoder := json.NewDecoder(r.Body)
	var beerData Beer
	err := decoder.Decode(&beerData)

	if(err != nil){
		w.WriteHeader(500)
	}
	defer r.Body.Close()
	beerData.ID = bson.NewObjectId()

	session.DB("prueba").C("beers").Insert(beerData)

	responseBeer(w,200, beerData)
}

func searchBeerById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	beerId := params["beerID"]

	session := mongoConn. Copy()
	defer session.Close()

	if !bson.IsObjectIdHex(beerId){
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(beerId)

	beer := Beer{}
	err := session.DB("prueba").C("beers").FindId(oid).One(&beer)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	responseBeer(w,200, beer)

}

func boxBeerPriceById(w http.ResponseWriter, r *http.Request){

	params := mux.Vars(r)
	beerId := params["beerID"]
	quantity := params["quantity"]
	currency := params["currency"]

	session := mongoConn. Copy()
	defer session.Close()

	if !bson.IsObjectIdHex(beerId){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(beerId)

	beer := Beer{}

	err := session.DB("prueba").C("beers").FindId(oid).One(&beer)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	priceBeerBox := PriceBeerBox{}

	change := changeCurrency(beer.Currency,currency,beer.Price)

	quantityBox, _ :=  strconv.ParseFloat(quantity, 64)

	if quantityBox <= 0 {
		quantityBox = 6
	}

	priceBeerBox.PriceTotal = quantityBox * change

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(priceBeerBox)

}

