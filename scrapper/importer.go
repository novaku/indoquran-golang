package scrapper

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"bitbucket.org/indoquran-api/helpers/logger"
	"bitbucket.org/indoquran-api/models"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2/bson"
)

const (
	importCSVFileLogTag = "scrapper/importer.go/ImportCSVFile"
)

// ImportCSVFile : import from CSV file to mongodb
func ImportCSVFile() {
	requestID := uuid.NewV4().String()
	lang := os.Getenv("LANG") // import language
	filePath := ""
	if lang == "en" {
		filePath = "./resources/English-Ahmed-Ali-100.csv"
	}
	if lang == "id" {
		filePath = "./resources/Indonesian-Bahasa-Indonesia-68.csv"
	}
	if lang == "ar" {
		filePath = "./resources/Arabic-(Original-Book)-1.csv"
	}

	logger.Info(importCSVFileLogTag, requestID, "import language: %s, file %s", lang, filePath)

	collection := models.DBConnect.MongoUse(models.DatabaseName, models.CollAyat)

	defer collection.Database.Session.Close()

	file, err := os.Open(filePath)

	if err != nil {
		logger.Error(importCSVFileLogTag, requestID, "Open file Error: %+v", err)
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true

	header := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Error(importCSVFileLogTag, requestID, "Read file Error: %+v", err)
			panic(err)
		}

		if !header {
			suratID, _ := strconv.Atoi(record[1])
			ayatID, _ := strconv.Atoi(record[2])

			selector := bson.M{
				"surat_id": suratID,
				"ayat_id":  ayatID,
			}

			update := bson.M{"$set": bson.M{"translate_" + lang: record[3]}}
			_, err = collection.Upsert(selector, update)
			if err != nil {
				logger.Error(importCSVFileLogTag, requestID, "UpdateDB Error: %+v", err)
				panic(err)
			}

			logger.Info(importCSVFileLogTag, requestID, "surat: %d, ayat: %d, lang: %s", suratID, ayatID, lang)
		}
		header = false
	}
}
