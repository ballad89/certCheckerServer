package handlers

import (
	"certcheckerServer/hub"
	"certcheckerServer/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Env struct {
	Db *models.DatabaseManager
}

func (e *Env) GetCerticates(w http.ResponseWriter, r *http.Request) {
	certs, err := e.Db.GetCertificates()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.MarshalIndent(certs, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (e *Env) GetCertificate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	certId := vars["id"]

	if !bson.IsObjectIdHex(certId) {
		http.NotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(certId)

	cert, err := e.Db.GetCertificateById(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	js, err := json.MarshalIndent(cert, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (e *Env) PostCertificate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := schema.NewDecoder()

	decoder.IgnoreUnknownKeys(true)

	dst := &models.Certificate{}

	err = decoder.Decode(dst, r.PostForm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	serverName := r.PostForm.Get("servername")

	if serverName == "" {
		http.Error(w, "servername field missing", http.StatusBadRequest)
		return
	}

	err = e.Db.InsertServerCertificate(dst, serverName)

	if err != nil {

		http.Error(w, err.Error(), http.StatusConflict)
		return

	}

	cert, err := e.Db.GetCertificateByCnameAndServer(dst.Cname, serverName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.MarshalIndent(cert, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hub.Publish(string(js))

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
