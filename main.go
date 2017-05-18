package main

import (
	"bitbucket.org/zetxek/molendb/molenDB"
)

func main() {

	d := db.OpenDB(db.GetDBName())
	defer d.Close()

	// Dam Square
	db.ClosestMill(d, 52.3759976, 4.8264306)
}
