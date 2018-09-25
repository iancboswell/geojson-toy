package main

import (
	"log"

	"github.com/iancboswell/geojson-toy/db"
	"github.com/iancboswell/geojson-toy/json"
)

func main() {
	// Insert & retrieve Geometry-containing struct
	p := db.InsertStructMember()

	// Create GeoJSON byte blob
	b := json.MarshalFeature(p)

	// Unmarshal GeoJSON byte blob & convert back to Pointy
	p2 := json.UnmarshalFeature(b)

	log.Printf("Unmarshalled Pointy Point: %v", p2.P.FlatCoords())

	db.InsertPoly()
}
