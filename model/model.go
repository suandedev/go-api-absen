package model

import "gorm.io/gorm"

type Absen struct {
	gorm.Model
	ID    int    `json:"id" gorm:"primaryKey autoIncrement"`
	Kelas string `json:"kelas"`
	Nama  string `json:"nama"`
	Waktu string `json:"waktu"`
	Ket   string `json:"ket"`
}

type Response struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Absen `json:"data"`
}
