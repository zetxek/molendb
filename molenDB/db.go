package db

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func GetDBName() string {
	return "./mills.db"
}

func OpenDB(name string) *sql.DB {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// ./files
func PopulateDB(dir string, db *sql.DB) {
	files, _ := ioutil.ReadDir(dir)

	for _, f := range files {
		fmt.Println(f.Name())
		mills := parseFile(fmt.Sprintf("./files/%s", f.Name()))
		for _, mill := range mills {
			Insert(db, mill)
		}
	}
}

type XMLMills struct {
	XMLName xml.Name `xml:"markers"`
	Mills   []Mill   `xml:"marker"`
}
type Mill struct {
	XMLName  xml.Name `xml:"marker"`
	Number   string   `xml:"nummer,attr"`
	Name     string   `xml:"name,attr"`
	Address  string   `xml:"address,attr"`
	Lat      float32  `xml:"lat,attr"`
	Lng      float32  `xml:"lng,attr"`
	MillType string   `xml:"type,attr"`
	Photo    string   `xml:"foto1,attr"`
}

func parseFile(file string) []Mill {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()

	var x XMLMills
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error unmarshalling: ", err)
	} else {
		xml.Unmarshal(data, &x)
		return x.Mills
	}

	return []Mill{}
}

func Createdb() {
	os.Remove("./mills.db")

	db, err := sql.Open("sqlite3", "./mills.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table mills (id integer not null primary key, name text, address text, lat float, lng float, type text, photo text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

func Insert(db *sql.DB, m Mill) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into mills(id, name, address, lat, lng, type, photo) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(m.Number, m.Name, m.Address, m.Lat, m.Lng, m.MillType, m.Photo)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

}

func ClosestMill(db *sql.DB, lat float32, lng float32) {

	stmt, err := db.Prepare("SELECT *, (((lat-?)*(lat-?)) + ((lng - ?)*(lng - ?)) ) *10000  as distance FROM mills ORDER BY distance ASC limit 1")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(lat, lat, lng, lng)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var address string
		var lat float32
		var lng float32
		var millType string
		var photo string
		var distance float32
		err = rows.Scan(&id, &name, &address, &lat, &lng, &millType, &photo, &distance)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, distance)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func ListItems(db *sql.DB) {

	rows, err := db.Query("select * from mills")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var address string
		var lat float32
		var lng float32
		var millType string
		var photo string
		err = rows.Scan(&id, &name, &address, &lat, &lng, &millType, &photo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
