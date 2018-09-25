package model

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

type Poly struct {
	P wkb.Polygon `db:"ST_AsWKB(poly)"`
}

func NewPoly(coords [][]geom.Coord) Poly {
	return Poly{wkb.Polygon{geom.NewPolygon(geom.XY).MustSetCoords(coords)}}
}
