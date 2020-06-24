package models

// UserModel defines the user model structure
type UserModel struct {
	Name       string `bson:"name"`
	Email      string `bson:"email"`
	Password   string `bson:"password"`
	IsVerified bool   `bson:"is_verified"`
}

// AyatModel defines the ayat model structure
type AyatModel struct {
	SuratID      int      `bson:"surat_id"`
	AyatID       int      `bson:"ayat_id"`
	BacalahID    string   `bson:"bacalah_id"`
	Read         string   `bson:"read"`
	TextIndo     string   `bson:"text_indo"`
	Penjelasan   []string `bson:"penjelasan"`
	Tafsir       string   `bson:"tafsir"`
	AsbabunNuzul string   `bson:"asbabun_nuzul"`
}
