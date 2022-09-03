package main

import (
	"encoding/json"
	"go-api-absen/middleware"
	"go-api-absen/model"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Absen struct {
	ID    int    `json:"id" gorm:"primaryKey autoIncrement"`
	Kelas string `json:"kelas"`
	Nama  string `json:"nama"`
	Waktu string `json:"waktu"`
	Ket   string `json:"ket"`
}

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("absenDb.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Absen{})
	return db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/absen", middleware.StoreAbsen).Methods("POST")
	r.HandleFunc("/absen", middleware.GetAbsens).Methods("GET")
	r.HandleFunc("/absen/{id}", middleware.GetAbsen).Methods("GET")
	r.HandleFunc("/absen/{id}", middleware.UpdateAbsen).Methods("PUT")
	r.HandleFunc("/absen/{id}", middleware.DeleteAbsen).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

func GetAbsen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := connectDB()

	var absen Absen
	db.First(&absen, params["id"])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(absen)
}
