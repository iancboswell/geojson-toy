package model

import (
	"log"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
	"github.com/twpayne/go-geom/encoding/wkb"
)

type Poly struct {
	P wkb.Polygon `db:"ST_AsWKB(poly)"`
}

func NewPoly(coords [][]geom.Coord) Poly {
	return Poly{wkb.Polygon{geom.NewPolygon(geom.XY).MustSetCoords(coords)}}
}

// Marshal the polygon itself as a GeoJSON.
func (p *Poly) Marshal() []byte {
	b, err := geojson.Marshal(p.P.Polygon)
	if err != nil {
		log.Fatalf("Failed to marshal polygon as GeoJSON: %v", err)
	}

	log.Printf("Marshalled the polygon on its own:\n%v\n", string(b))

	return b
}

func UnmarshalPolygon(b []byte) Poly {
	return Poly{} // TODO
}

// Marshal the polygon as part of a GeoJSON Feature.
func (p *Poly) MarshalAsFeature() []byte {
	f := geojson.Feature{Geometry: p.P.Polygon}
	b, err := f.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature: %v", err)
	}

	log.Printf("Marshalled the polygon as part of a feature:\n%v\n", string(b))

	return b
}

// Turn verbosity up to 11 with a FeatureCollection(tm)!
func (p *Poly) MarshalAsFeatureCollection() []byte {
	f := geojson.Feature{Geometry: p.P.Polygon}
	fc := geojson.FeatureCollection{[]*geojson.Feature{&f}}
	b, err := fc.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature collection: %v", err)
	}

	log.Printf("Marshalled the polygon as part of a feature collection:\n%v\n", string(b))

	return b
}
