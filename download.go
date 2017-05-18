package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func dlFile(url, filename string) {
	fmt.Println("Downloading " + url + " ...")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)
}

func Download() {

	var base = "https://molendatabase.nl/nederland/bewaartypemarkers.php?type="
	var types = [...]string{
		"Beltmolen",
		"Grondzeiler",
		"Tjasker",
		"Kleine molen",
		"Paltrokmolen",
		"Spinnenkop",
		"Stellingmolen",
		"Wipmolen",
		"Weidemolen",
		"Standerdmolen",
		"Gesloten standerdmolen",
		"Torenmolen",
		"Tonmolen",
		"Tredmolen",
		"Achtkante molen",
		"Zeskante molen",
		"Achtkante binnenkruier",
		"Watermolen"}

	for _, element := range types {
		var f = strings.ToLower(element)

		f = strings.Replace(f, " ", "", -1)
		f = fmt.Sprintf("%s.xml", f)
		var url = fmt.Sprintf("%s%s", base, element)

		dlFile(url, f)
		fmt.Println(f + " saved!")
	}

}
