package main

import (
	"encoding/json"
	"log"

	"github.com/iancboswell/geojson-toy/db"
	"github.com/iancboswell/geojson-toy/model"
)

func main() {
	p := db.InsertPointy()

	j, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	log.Printf("Marshalled JSON:\n%s\n", string(j))

	log.Printf("Manual Marshal:\n%s\n", string(p.Marshal()))
}

func Munge() {
	// Insert & retrieve Point-containing struct
	p := db.InsertPointy()

	// Create GeoJSON byte blob
	b := p.Marshal()

	// Unmarshal GeoJSON byte blob & convert back to Pointy
	p2 := model.UnmarshalPoint(b)

	log.Printf("Unmarshalled Pointy Point: %v", p2.P.FlatCoords())

	// Insert & retrieve Polygon-containing struct
	po := db.InsertPoly()

	// Create GeoJSON byte blob
	bo := po.Marshal()

	// Unmarshal GeoJSON byte blob & convert back to Pointy
	po2 := model.UnmarshalPolygon(bo)

	log.Printf("Unmarshalled Poly Polygon: %v", po2.P.FlatCoords())
}
