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

	return b
}

// Unmarshal a GeoJSON byte blob into a Pointy struct.
func UnmarshalPoint(b []byte) Pointy {
	var g geom.T = geom.NewPoint(geom.XY)
	err := geojson.Unmarshal(b, &g)
	if err != nil {
		log.Fatalf("Failed to unmarshal point: %v", err)
	}

	return Pointy{wkb.Point{g.(*geom.Point)}}
}

// Marshal the point as part of a GeoJSON Feature.
func (p *Pointy) MarshalAsFeature() []byte {
	f := geojson.Feature{Geometry: p.P.Point}
	b, err := f.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature: %v", err)
	}

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
	return Pointy{wkb.Point{g.(*geom.Point)}}
}

// Turn verbosity up to 11 with a FeatureCollection(tm)!
func (p *Pointy) MarshalAsFeatureCollection() []byte {
	f := geojson.Feature{Geometry: p.P.Point}
	fc := geojson.FeatureCollection{[]*geojson.Feature{&f}}
	b, err := fc.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal GeoJSON feature collection: %v", err)
	}

	return b
}
