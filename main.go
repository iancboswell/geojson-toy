package main

import (
	"log"

	"github.com/iancboswell/geojson-toy/db"
	"github.com/iancboswell/geojson-toy/model"
)

func main() {
	// Insert & retrieve Point-containing struct
	p := db.InsertStructMember()

	// Create GeoJSON byte blob
	b := p.MarshalAsFeature()

	// Unmarshal GeoJSON byte blob & convert back to Pointy
	p2 := model.UnmarshalPointAsFeature(b)

	log.Printf("Unmarshalled Pointy Point: %v", p2.P.FlatCoords())

	// Insert & retrieve Polygon-containing struct
	po := db.InsertPoly()

	// Create GeoJSON byte blob
	bo := po.Marshal()

	// Unmarshal GeoJSON byte blob & convert back to Pointy
	po2 := model.UnmarshalPolygon(bo)

	log.Printf("Unmarshalled Poly Polygon: %v", po2.P.FlatCoords())
}
