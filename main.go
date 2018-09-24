package main

import (
	"fmt"
	"log"

	// Make sure the MySQL driver has called `init()`
	_ "github.com/go-sql-driver/mysql"

	//sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

func main() {
	db, err := sqlx.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&multiStatements=true&readTimeout=1s&clientFoundRows=true",
		"root",
		"root",
		"(127.0.0.1:3306)",
		"geojson_toy"))
	if err != nil {
		log.Fatalf("SQLX failed to open a connection to MySQL\n"+
			"sqlx.open() failed with: %v", err)
	}

	// Define a point in 2-D space.
	p := wkb.Point{
		Point: geom.NewPoint(geom.XY).
			MustSetCoords(
				geom.Coord{-77.8187259, 40.8089934})}

	log.Printf("Inserting point with coordinates:\n(%v, %v)\n", p.FlatCoords()[0], p.FlatCoords()[1])

	result, err := db.Exec(`INSERT INTO points (pt) VALUES (ST_GeomFromWKB(?));`, &p)
	if err != nil {
		log.Fatalf("SQL INSERT failed: %v", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Introspection failed: %v", err)
	}

	ret := wkb.Point{Point: geom.NewPoint(geom.XY)}
	err = db.Get(&ret, "SELECT ST_AsWKB(pt) FROM points WHERE id = ?", lastID)
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}

	log.Printf("Retrieved point with coordinates:\n(%v, %v)\n", ret.FlatCoords()[0], ret.FlatCoords()[1])

	// Generate SQL insert statement.
	/*query, args, err := sq.Insert("").Into("points").Columns("pt").
		Values(&p).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Fatalf("SQL execution error: %v", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("SQL introspection failed: %v", err)
	}

	log.Printf("New row ID: %v", lastID)*/
}
