package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dmitryk-dk/form_phone/server/database"
)

type PhoneNumber struct {
	Number string `json:"phoneNumber"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var phoneNumber *PhoneNumber
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("parsing error: %v\n", err)
		w.Write([]byte("error"))
	}
	json.Unmarshal(body, &phoneNumber)
	fmt.Println(phoneNumber)
	w.Write([]byte("success"))
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var dbHelper database.DbMethodsHelper
	dbHelper = &database.DbMethods{}
	phones, err := dbHelper.GetPhones()
	jsonPhones, err := json.Marshal(phones)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonPhones)
}
