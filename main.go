package main

import (
	"bitbucket.org/zetxek/molendb/molenDB"
)

func main() {

	d := db.OpenDB(db.GetDBName())
	defer d.Close()

	//Van bossestraat
	db.ClosestMill(d, 52.3759129, 4.8643734)
	// Jan Haringstraat
	db.ClosestMill(d, 52.3760028, 4.8603556)
}
