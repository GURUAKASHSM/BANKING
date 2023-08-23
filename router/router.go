package router

import (
	"mongoapi/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	routes:= mux.NewRouter()
    routes.HandleFunc("/get",controllers.Getalldata).Methods("GET")
	routes.HandleFunc("/getid/{id}",controllers.Getdatabyid).Methods("GET")
	routes.HandleFunc("/update/{id}",controllers.UpdateProfile).Methods("PUT")
	routes.HandleFunc("/insert",controllers.CreateProfile).Methods("POST")
	routes.HandleFunc("/delete/{id}",controllers.Deleteprofile).Methods("DELETE")
	routes.HandleFunc("/deleteall",controllers.Deleteallprofile).Methods("DELETE")
	return routes
}