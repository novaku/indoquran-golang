package scrapper

import (
	"fmt"
	"strconv"
	"sync"

	"indoquran-golang/models"
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
	defer collection.Database.Session.Close()

	model := &models.AyatModel{}

	idString := strconv.Itoa(id)
	ayatSurat := models.StaticAyatID[idString]
	model.BacalahID = idString
	model.AyatID = ayatSurat.AyatID
	model.SuratID = ayatSurat.SuratID

	model.Read, model.TextIndo, model.Penjelasan = AyatID(idString)
	model.Tafsir = TafsirID(idString)
	model.AsbabunNuzul = AsbabunNuzulID(idString)

	// fmt.Println(m)
	err := collection.Insert(model)
	if err != nil {
		fmt.Println("Error insert mongodb: ", err)
	}
	fmt.Println("AYAT_ID: ", id)
	fmt.Println("===========================================================================")
}
