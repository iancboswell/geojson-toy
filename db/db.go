package db

import (
	"fmt"
	"log"

	// Make sure the MySQL driver has called `init()`
	_ "github.com/go-sql-driver/mysql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"

	"github.com/iancboswell/geojson-toy/model"
)

var db *sqlx.DB

func init() {
	var err error

	db, err = sqlx.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&multiStatements=true&readTimeout=1s&clientFoundRows=true",
		"root",
		"root",
		"(127.0.0.1:3306)",
		"geojson_toy"))
	if err != nil {
		log.Fatalf("SQLX failed to open a connection to MySQL\n"+
			"sqlx.open() failed with: %v", err)
	}

}

func InsertPointy() model.Pointy {
	// Define a point in 2-D space as a struct member.
	p := model.Pointy{P: wkb.Point{
		Point: geom.NewPoint(geom.XY).
			MustSetCoords(
				geom.Coord{40.8089934, -77.8187259})}}

	log.Printf("Inserting point with coordinates:\n%v\n", p.P.FlatCoords())

	query, args, err := sq.Insert("").Into("points").Columns("pt").
		Values(sq.Expr("ST_GeomFromWKB(?)", &p.P)).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Fatalf("SQL insert execution error: %v", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("SQL introspection failed: %v", err)
	}

	var ret model.Pointy

	query, args, err = sq.Select("ST_AsWKB(pt)").From("points").
		Where(sq.Eq{"id": lastID}).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	err = db.Get(&ret, query, args...)
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}

	log.Printf("Retrieved point with coordinates:\n%v\n", ret.P.FlatCoords())

	return ret
}

func InsertPoly() model.Poly {
	coords := [][]geom.Coord{{
		{40.8089934, -77.8187259},
		{18.470299, -66.719073},
		{35.678105, -108.114275},
		{44.940218, -93.170220},
		{41.878113, -87.629799},
		{40.8089934, -77.8187259}}}

	p := model.NewPoly(coords)

	log.Printf("Inserting polygon with coordinates:\n%v\n", p.P.FlatCoords())

	query, args, err := sq.Insert("").Into("polygons").Columns("poly").
		Values(sq.Expr("ST_GeomFromWKB(?)", &p.P)).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Fatalf("SQL insert execution error: %v", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("SQL introspection failed: %v", err)
	}

	var ret model.Poly

	query, args, err = sq.Select("ST_AsWKB(poly)").From("polygons").
		Where(sq.Eq{"id": lastID}).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	err = db.Get(&ret, query, args...)
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}

	log.Printf("Retrieved polygon with coordinates:\n%v\n", ret.P.FlatCoords())

	return ret
}
