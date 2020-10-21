package scrapper

import (
	"fmt"
	"strconv"
	"sync"

	"bitbucket.org/indoquran-api/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	ayatStart = 1988
	ayatEnd   = 8223
)

// ScrapQuranBacalahNet : scrap http://quran.bacalah.net/content/surat
func ScrapQuranBacalahNet() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := ayatStart; i <= ayatEnd; i++ {
		wg.Add(1)
		go worker(i, &wg, &m)
	}

	wg.Wait()
}

func worker(id int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock

	collection := models.DBConnect.MongoUse(models.DatabaseName, models.CollAyat)

	defer wg.Done()
	defer m.Unlock()
	// defer collection.Database.Session.Close()

	model := &models.AyatModel{}

	idString := strconv.Itoa(id)
	ayatSurat := models.StaticAyatID[idString]
	model.BacalahID = idString
	model.AyatID = ayatSurat.AyatID
	model.SuratID = ayatSurat.SuratID

	model.Read, model.TextIndo, model.Penjelasan = AyatID(idString)
	model.Tafsir = TafsirID(idString)
	model.AsbabunNuzul = AsbabunNuzulID(idString)

	selector := bson.M{
		"surat_id": ayatSurat.SuratID,
		"ayat_id":  ayatSurat.AyatID,
	}

	_, err := collection.Upsert(selector, model)

	// err := collection.Insert(model)
	if err != nil {
		fmt.Printf("Error upsert mongodb: %+v", err)
	}
	fmt.Printf("surat: %d, ayat: %d, ayat_id: %d\n", ayatSurat.SuratID, ayatSurat.AyatID, id)
	fmt.Println("===========================================================================")
}
