package json

import (
	"log"

	"github.com/twpayne/go-geom/encoding/geojson"

	"github.com/iancboswell/geojson-toy/model"
)

// Marshal the point itself as a GeoJSON.
func MarshalPoint(p model.Pointy) []byte {
	b, err := geojson.Marshal(p.P.Point)
	if err != nil {
		log.Fatalf("Failed to marshal point as GeoJSON: %v", err)
	}

	log.Printf("Marshalled the point on its own:\n%v\n", string(b))

	return b
}

// Marshal the point as part of a GeoJSON Feature.
func MarshalFeature(p model.Pointy) []byte {
	f := geojson.Feature{Geometry: p.P.Point}
	b, err := f.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature: %v", err)
	}

	log.Printf("Marshalled the point as part of a feature:\n%v\n", string(b))

	return b
}

// Unmarshal the GeoJSON Feature into a Pointy struct.
func UnmarshalFeature(b []byte) model.Pointy {
	var f geojson.Feature
	err := f.UnmarshalJSON(b)
	if err != nil {
		log.Fatalf("Failed to unmarshal feature: %v", err)
	}

	g := f.Geometry
	return model.NewPointy(g.FlatCoords())
}

// Turn verbosity up to 11 with a FeatureCollection(tm)!
func MarshalFeatureCollection(p model.Pointy) []byte {
	f := geojson.Feature{Geometry: p.P.Point}
	fc := geojson.FeatureCollection{[]*geojson.Feature{&f}}
	b, err := fc.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature collection: %v", err)
	}

	log.Printf("Marshalled the point as part of a feature collection:\n%v\n", string(b))

	return b
}
