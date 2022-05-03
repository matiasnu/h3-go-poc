/*
* Para la creacion de esta poc me baso en el notebook https://observablehq.com/@nrabinowitz/h3-radius-lookup
 */
package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/uber/h3-go/v3"
)

type ReferencePoint struct {
	ID      string `bson:"-"`
	Indexh3 string `bson:"indexh3"`
	ShopID  string `bson:"shopid"`
}

type Shop struct {
	ID        string
	Latitude  float64
	Longitude float64
}

// La resolucion se basa en cuan mas precisos van a ser los hexagonos
// cuanto mayor sea la resolucion mas hexagonos van a crearse por km, y a su
// vez eso consume mas memoria y mas division.
const resolution = 8

// La siguiente constante indica el radio en Kilometros en el que queremos hacer la busqueda
const searchRadiusKm = 5

func main() {
	ExampleFromGeo()
}

func ExampleFromGeo() {
	geo := h3.GeoCoord{
		Latitude:  -34.65368544117287, // Random Cooordinate donde se encuentra el user
		Longitude: -58.513967653931594,
	}
	initDb()
	initDataRandomDB()
	kRingResults(geo)
}

func convertGeoCordToIndex(geo h3.GeoCoord) h3.H3Index {
	return h3.FromGeo(geo, resolution)
}

func convertIndexesArrayToStringArray(lookupIndexes []h3.H3Index) []string {
	var list []string
	for _, i := range lookupIndexes {
		list = append(list, h3.ToString(i))
	}
	return list
}

// geo es la coordenada donde se encuentra el usuario
// searchRadiusKm es el radio en km que queremos buscar al rededor
// devuelve un ring o anillo dentro de las coordenadas buscadas
// donde se encuentran los demas vecinos
func kRingIndexes(origin h3.H3Index, searchRadiusKm float64) []h3.H3Index {

	// Transform the radius from km to grid distance
	radius := math.Floor(searchRadiusKm / (h3.EdgeLengthKm(resolution) * 2))
	return h3.KRing(origin, int(radius))
}

func kRingResults(geo h3.GeoCoord) {
	origin := convertGeoCordToIndex(geo)
	lookupIndexes := kRingIndexes(origin, searchRadiusKm)
	// Find all points of interest in the k-ring
	fmt.Printf("%#x\n", lookupIndexes)
	indexesToString := convertIndexesArrayToStringArray(lookupIndexes)
	getShopsByRing(indexesToString)
}

// Esta funcion permite agregar shops a la base de datos de forma que calcula el indice que pertene al hexagono
// y lo agrega luego a la collection shops.
func AddShopByIndex(shop Shop) {
	geo := h3.GeoCoord{
		Latitude:  shop.Latitude,
		Longitude: shop.Longitude,
	}
	index := convertGeoCordToIndex(geo)
	addShop(shop.ID, h3.ToString(index))
}

// Genera shops random en la base de datos para poder realizar la busqueda
func initDataRandomDB() {
	shopsCreated := 100
	for i := 0; i < shopsCreated; i++ {
		location := RandLocation(Location{
			Latitude:  -34.65368544117287,
			Longitude: -58.513967653931594,
		}, 1)
		shopTest := Shop{
			ID:        strconv.Itoa(i),
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		}
		AddShopByIndex(shopTest)
	}
}
