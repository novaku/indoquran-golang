package scrapper

import (
	"bufio"
	"encoding/csv"
	"indoquran-golang/models"
	"io"
	"os"
	"strconv"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

// ImportCSVFile : import from CSV file to mongodb
func ImportCSVFile(path, lang string) {
	glog.Infof("import language: %s, file %s", lang, path)

	collection := models.DBConnect.MGOUse(models.DatabaseName, models.CollAyat)

	defer collection.Database.Session.Close()

	file, err := os.Open(path)

	if err != nil {
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
			err = collection.Update(selector, update)
			if err != nil {
				glog.Errorf("Error updating surat: %d, ayat: %d, error: %+v", suratID, ayatID, err)
				panic(err)
			}

			glog.Infof("Updating surat: %d, ayat: %d, lang: %s", suratID, ayatID, lang)
		}
		header = false
	}
}
