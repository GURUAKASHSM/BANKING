package controllers

import (
	"encoding/json"
	"mongoapi/models"
	"mongoapi/service"
	"net/http"
	"github.com/gorilla/mux"
)

func Getalldata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := service.Getalldata()
	json.NewEncoder(w).Encode(allmovies)
}
func Getdatabyid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control Allow-Methods", "POST")
	params := mux.Vars(r)
	data := service.Getdatabyid(params["id"])
	json.NewEncoder(w).Encode(data)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var profile models.Profile
	json.NewDecoder(r.Body).Decode(&profile)
	service.Insert(profile)
	json.NewEncoder(w).Encode(profile)
}
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control Allow-Methods", "PUT")
	params := mux.Vars(r)
	service.Updateone(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func Deleteprofile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Controls Allow-Methods", "DELETE")
	params := mux.Vars(r)
	service.DeleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func Deleteallprofile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control Allow-Methods", "DELETE")
	count := service.DeleteAll()
	json.NewEncoder(w).Encode(count)
}
