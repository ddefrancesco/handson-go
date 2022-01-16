package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ddefrancesco/handson-go/jsonstore/helper"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	postBody, _ := ioutil.ReadAll(r.Body)
	Package.Data = string(postBody)
	driver.db.Save(&Package)
	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)
	driver.db.First(&Package, vars["id"])
	var DatiPackage interface{}

	json.Unmarshal([]byte(Package.Data), &DatiPackage)
	var response = PackageResponse{Package: Package}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(response)
	w.Write(respJSON)
}

func (driver *DBClient) PostShipment(w http.ResponseWriter, r *http.Request) {
	var Shipment = helper.Shipment{}
	postBody, _ := ioutil.ReadAll(r.Body)
	Shipment.Data = string(postBody)
	driver.db.Save(&Shipment)
	responseMap := map[string]interface{}{"id": Shipment.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

func main() {
	db, _ := helper.InitDB()
	dbclient := &DBClient{db: db}
	r := mux.NewRouter()
	r.HandleFunc("/v1/package", dbclient.PostPackage).Methods("POST")
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}", dbclient.GetPackage).Methods("GET")
	r.HandleFunc("/v1/shipment", dbclient.PostShipment).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Connected to server %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
