package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func (e *Env) GetServers(w http.ResponseWriter, r *http.Request) {
	servs, err := e.Db.GetServers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.MarshalIndent(servs, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (e *Env) GetServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	servId := vars["id"]

	if !bson.IsObjectIdHex(servId) {
		http.NotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(servId)

	serv, err := e.Db.GetServerById(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	js, err := json.MarshalIndent(serv, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (e *Env) GetServerCertificates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	servId := vars["id"]

	if !bson.IsObjectIdHex(servId) {
		http.NotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(servId)

	servers, err := e.Db.GetServerCertificates(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	js, err := json.MarshalIndent(servers, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
