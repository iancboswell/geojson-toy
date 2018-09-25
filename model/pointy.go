package model

import (
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

type Pointy struct {
	P wkb.Point `db:"ST_AsWKB(pt)"`
}

func NewPointy(coords []float64) Pointy {
	return Pointy{wkb.Point{geom.NewPoint(geom.XY).
		MustSetCoords(coords)}}
}
