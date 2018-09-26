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

	return b
}

// Unmarshal a GeoJSON byte blob into a Poly struct.
func UnmarshalPolygon(b []byte) Poly {
	var g geom.T = geom.NewPolygon(geom.XY)
	err := geojson.Unmarshal(b, &g)
	if err != nil {
		log.Fatalf("Failed to unmarshal polygon: %v", err)
	}

	return Poly{wkb.Polygon{g.(*geom.Polygon)}}
}

// Marshal the polygon as part of a GeoJSON Feature.
func (p *Poly) MarshalAsFeature() []byte {
	f := geojson.Feature{Geometry: p.P.Polygon}
	b, err := f.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature: %v", err)
	}

	return b
}

// Unmarshal the GeoJSON Feature into a Poly struct.
func UnmarshalPolygonAsFeature(b []byte) Poly {
	var f geojson.Feature
	err := f.UnmarshalJSON(b)
	if err != nil {
		log.Fatalf("Failed to unmarshal feature: %v", err)
	}

	g := f.Geometry
	return Poly{wkb.Polygon{g.(*geom.Polygon)}}
}

// Turn verbosity up to 11 with a FeatureCollection(tm)!
func (p *Poly) MarshalAsFeatureCollection() []byte {
	f := geojson.Feature{Geometry: p.P.Polygon}
	fc := geojson.FeatureCollection{[]*geojson.Feature{&f}}
	b, err := fc.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature collection: %v", err)
	}

	return b
}
