package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"go-api-absen/fun"
	"go-api-absen/model"

	"github.com/gorilla/mux"
)

func StoreAbsen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get data student
	kelas := "XII RPL 1"
	nama := "Rizky"
	waktu := time.Now().Format("2006-01-02 15:04:05")
	ket := "Hadir"

	// store data student
	res := fun.StoreAbsen(kelas, nama, waktu, ket)

	if res.Error != nil {
		panic(res.Error)
	}

	// response
	var absen model.Absen
	res.Scan(&absen)
	json.NewEncoder(w).Encode(absen)
}

func GetAbsens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get data student
	absens := fun.GetAbsens()

	// response
	json.NewEncoder(w).Encode(absens)
}

func GetAbsen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// get data student
	absen := fun.GetAbsen(id)

	// // response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(absen)
}

func UpdateAbsen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get id
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])

	// get data student
	kelas := "XII RPL 1"
	nama := "Rizky"
	waktu := time.Now().Format("2006-01-02 15:04:05")
	ket := "alfa"

	// update data student
	res := fun.UpdateAbsen(id, kelas, nama, waktu, ket)

	// response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteAbsen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get id
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	// delete data student
	res := fun.DeleteAbsen(id)

	if res.Error != nil {
		panic(res.Error)
	}

	respon := model.Response{
		Status:  http.StatusOK,
		Message: "data deleted",
		Data:    nil,
	}

	// response
	json.NewEncoder(w).Encode(respon)
}
