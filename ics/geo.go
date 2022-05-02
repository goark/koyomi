package ics

import "database/sql"

type Geo struct {
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
}

/** These codes are forked form "github.com/PuloV/ics-golang" package. (licensed under MIT) */
