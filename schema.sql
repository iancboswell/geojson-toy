# Create a `geojson_toy` database and run this script against it.

CREATE TABLE IF NOT EXISTS points (
    id          INT         AUTO_INCREMENT  NOT NULL, PRIMARY KEY(id),
    created_at  DATETIME    NOT NULL        DEFAULT NOW(),
    updated_at  DATETIME    NOT NULL        DEFAULT NOW() ON UPDATE NOW(),

    pt          POINT       NOT NULL
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS polygons (
    id          INT         AUTO_INCREMENT  NOT NULL, PRIMARY KEY(id),
    created_at  DATETIME    NOT NULL        DEFAULT NOW(),
    updated_at  DATETIME    NOT NULL        DEFAULT NOW() ON UPDATE NOW(),

    poly        POLYGON     NOT NULL
) ENGINE=InnoDB;