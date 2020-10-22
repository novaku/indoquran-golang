package models

import "gopkg.in/mgo.v2/bson"

// UserModel defines the user model structure
type UserModel struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Email      string        `bson:"email"`
	Password   string        `bson:"password"`
	IsVerified bool          `bson:"is_verified"`
}

// AyatModel defines the ayat model structure
type AyatModel struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	SuratID      int           `bson:"surat_id"`
	AyatID       int           `bson:"ayat_id"`
	BacalahID    string        `bson:"bacalah_id"`
	Read         string        `bson:"read"`
	TextIndo     string        `bson:"text_indo"`
	TranslateAR  string        `bson:"translate_ar"`
	TranslateID  string        `bson:"translate_id"`
	TranslateEN  string        `bson:"translate_en"`
	Penjelasan   []string      `bson:"penjelasan"`
	Tafsir       string        `bson:"tafsir"`
	AsbabunNuzul string        `bson:"asbabun_nuzul"`
}
