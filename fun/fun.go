package fun

import (
	"go-api-absen/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("absenDb.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Absen{})
	return db
}

func StoreAbsen(kelas string, nama string, waktu string, ket string) *gorm.DB {
	db := connectDb()

	absen := model.Absen{
		Kelas: kelas,
		Nama:  nama,
		Waktu: waktu,
		Ket:   ket,
	}
	res := db.Create(&absen)

	if res.Error != nil {
		panic("failed to store data")
	}

	return res
}

func GetAbsens() []model.Absen {
	db := connectDb()

	var absens []model.Absen
	db.Find(&absens)

	return absens
}

func GetAbsen(id int) model.Absen {
	db := connectDb()

	var absen model.Absen
	db.First(&absen, id)

	return absen
}

func UpdateAbsen(id int, kelas string, nama string, waktu string, ket string) model.Absen {
	db := connectDb()

	var absen model.Absen
	db.Model(&absen).Where("id = ?", id).Updates(model.Absen{Kelas: kelas, Nama: nama, Waktu: waktu, Ket: ket})

	return absen
}

func DeleteAbsen(id int) *gorm.DB {
	db := connectDb()

	res := db.Delete(&model.Absen{}, id)

	return res
}
