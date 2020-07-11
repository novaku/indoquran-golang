package main

import (
	"flag"
	"indoquran-golang/models"
	"indoquran-golang/scrapper"
	"indoquran-golang/server"
	"os"
)

const (
	yes = "1"
)

// init gets called before the main function
func init() {
	flag.Parse()
	scrap := os.Getenv("SCRAPP")    // to scrap from web
	importDB := os.Getenv("IMPORT") // to import from csv to mongodb
	if scrap == yes {
		models.InitializeStaticAyatSuratID()
		scrapper.ScrapQuranBacalahNet() // do scrapping to http://quran.bacalah.net/content/surat/GetContentAyat.php
	}
	if importDB == yes {
		scrapper.ImportCSVFile()
	}
}

// main function
func main() {
	server.StartServer()
}
