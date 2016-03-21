package models

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestConnection(t *testing.T) {
	dM, err := setup()

	if err != nil {
		t.Errorf("connection failed")
	}

	dM.Close()

}

func TestInsertGetAndDeleteCertificate(t *testing.T) {
	dM, err := setup()

	if err != nil {
		t.Errorf("connection failed")
	}
	defer dM.Close()

	id := bson.NewObjectId()

	cert := Certificate{
		ID:               id,
		Cname:            "test",
		SigningAlgorithm: "sha256",
		Issuer:           "me",
		NotAfter:         "25-02-2016",
		NotBefore:        "25-02-2014",
	}
	err = dM.InsertServerCertificate(&cert, "server1")

	if err != nil {
		t.Error(err)
	}

	ce, err := dM.GetCertificateById(id)

	if err != nil {
		t.Error(err)
	}

	if ce.Cname != "test" {

		t.Error("Cname not properly set")
	}

	if ce.SigningAlgorithm != "sha256" {
		t.Error("signing algorithm not properly set")
	}

	certs, err := dM.GetCertificates()

	if len(certs) != 1 {
		t.Log(len(certs))
		t.Error("Wrong number of certs")
	}

	ce1, err := dM.GetCertificateByCnameAndServer("test", "server1")

	if err != nil {
		t.Error(err)
	}

	if ce1.Issuer != "me" {
		t.Error("issuer field does not match")
	}

	certs1, err := dM.GetServerCertificates(ce1.ServerId)

	if err != nil {
		t.Error(err)
	}

	if len(certs1) == 0 {
		t.Error("Certs not found")
	}

	err = dM.DeleteServerById(ce1.ServerId)

	if err != nil {
		t.Error(err)
	}

	err = dM.DeleteCertificateById(id)

	if err != nil {
		t.Error(err)
	}

}

func TestInsertGetAndDeleteServer(t *testing.T) {
	dM, err := setup()

	if err != nil {
		t.Errorf("connection failed")
	}
	defer dM.Close()

	id := bson.NewObjectId()

	serv := Server{
		ID:   id,
		Name: "testServer",
	}
	err = dM.InsertServer(&serv)

	if err != nil {
		t.Error(err)
	}

	s, err := dM.GetServerById(id)

	if err != nil {
		t.Error(err)
	}

	if s.Name != "testServer" {

		t.Error("Server name not properly set")
	}

	s1, err := dM.GetServerByName("testServer")

	if err != nil {
		t.Error(err)
	}

	if s1.ID != id {
		t.Error("Server not returned by name")
	}

	servs, err := dM.GetServers()

	if len(servs) != 1 {
		t.Log(len(servs))
		t.Error("Wrong number of servers")
	}

	err = dM.DeleteServerById(id)

	if err != nil {
		t.Error(err)
	}

	_, err = dM.GetServerByName("testServer")

	if err == nil {
		t.Error("object found when it does not exist")
	}

}

func setup() (*DatabaseManager, error) {
	dM, err := ConnectToDatabase("mongodb://33.33.13.45", "testdb1")

	return dM, err
}
