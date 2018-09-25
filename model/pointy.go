package model

import (
	"log"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/twpayne/go-geom/encoding/wkb"
)

type Pointy struct {
	P wkb.Point `db:"ST_AsWKB(pt)"`
}

func NewPointy(coords []float64) Pointy {
	return Pointy{wkb.Point{geom.NewPoint(geom.XY).
		MustSetCoords(coords)}}
}

// Marshal the point itself as a GeoJSON.
func (p *Pointy) Marshal() []byte {
	b, err := geojson.Marshal(p.P.Point)
	if err != nil {
		log.Fatalf("Failed to marshal point as GeoJSON: %v", err)
	}

	log.Printf("Marshalled the point on its own:\n%v\n", string(b))

	return b
}

// Marshal the point as part of a GeoJSON Feature.
func (p *Pointy) MarshalAsFeature() []byte {
	f := geojson.Feature{Geometry: p.P.Point}
	b, err := f.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature: %v", err)
	}

	log.Printf("Marshalled the point as part of a feature:\n%v\n", string(b))

	return b
}

// Unmarshal the GeoJSON Feature into a Pointy struct.
func UnmarshalPointAsFeature(b []byte) Pointy {
	var f geojson.Feature
	err := f.UnmarshalJSON(b)
	if err != nil {
		log.Fatalf("Failed to unmarshal feature: %v", err)
	}

	g := f.Geometry
	return NewPointy(g.FlatCoords())
}

// Turn verbosity up to 11 with a FeatureCollection(tm)!
func (p *Pointy) MarshalAsFeatureCollection() []byte {
	f := geojson.Feature{Geometry: p.P.Point}
	fc := geojson.FeatureCollection{[]*geojson.Feature{&f}}
	b, err := fc.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature collection: %v", err)
	}

	log.Printf("Marshalled the point as part of a feature collection:\n%v\n", string(b))

	return b
}
