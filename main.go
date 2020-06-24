package main

import (
	"flag"
	"indoquran-golang/models"
	"indoquran-golang/scrapper"
	"indoquran-golang/server"
	"os"
)

// init gets called before the main function
func init() {
	flag.Parse()
	scrap := os.Getenv("SCRAPP")
	if scrap == "1" {
		models.InitializeStaticAyatSuratID()
		scrapper.ScrapQuranBacalahNet() // do scrapping to http://quran.bacalah.net/content/surat/GetContentAyat.php
	}
}

// main function
func main() {
	server.StartServer()
}
