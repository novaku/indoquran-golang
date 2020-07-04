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
	scrap := os.Getenv("SCRAPP")
	importDB := os.Getenv("IMPORT")
	if scrap == yes {
		models.InitializeStaticAyatSuratID()
		scrapper.ScrapQuranBacalahNet() // do scrapping to http://quran.bacalah.net/content/surat/GetContentAyat.php
	}
	if importDB == yes {
		lang := os.Getenv("LANG")
		filePath := ""
		if lang == "en" {
			filePath = "./resurces/English-Yusuf-Ali-59.csv"
		}
		if lang == "id" {
			filePath = "./resurces/Indonesian-Bahasa-Indonesia-68.csv"
		}
		if lang == "ar" {
			filePath = "./resurces/Arabic-(Original-Book)-1.csv"
		}
		scrapper.ImporrtCSVFile(filePath, lang)
	}
}

// main function
func main() {
	server.StartServer()
}
