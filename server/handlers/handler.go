package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dmitryk-dk/form_phone/server/database"
	"github.com/dmitryk-dk/form_phone/server/models"
    "github.com/dmitryk-dk/form_phone/server/config"
)

type PhoneNumber struct {
	Number string `json:"phoneNumber"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	dbHelper := &database.DbMethods{}
	phone := &models.Phone{}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		if err = json.Unmarshal(body, phone); err != nil {
			log.Println("Unmarshall error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Wrong data", http.StatusInternalServerError)
		}

		err = dbHelper.AddPhone(phone)
		if err != nil {
			log.Println("DB Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Can't connect to DataBase", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("success"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Used wrong method", http.StatusInternalServerError)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	dbHelper := &database.DbMethods{}

	if r.Method == "GET" {

		phones, err := dbHelper.GetPhones()
		if err != nil {
			log.Println("DB Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Can't connect to DataBase", http.StatusInternalServerError)
		}

		jsonPhones, err := json.Marshal(phones)
		if err != nil {
			log.Println("Marshal error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Wrong data", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonPhones)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Used wrong method", http.StatusInternalServerError)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	dbHelper := &database.DbMethods{}
	phone := &models.Phone{}

	if r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		if err = json.Unmarshal(body, phone); err != nil {
			log.Println("Unmarshall error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Wrong data", http.StatusInternalServerError)
		}

		err = dbHelper.DeletePhone(phone)
		if err != nil {
			log.Println("DB Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Can't connect to DataBase", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("success"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Used wrong method", http.StatusInternalServerError)
	}
}

func UiConfigHandler(w http.ResponseWriter, r *http.Request) {
	conf := config.GetUIConfig()
	content, err := json.Marshal(conf)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
        return
    }
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
