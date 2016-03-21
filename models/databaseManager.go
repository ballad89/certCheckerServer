package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	certificatesCollection = "certificates"
	ServersCollection      = "servers"
)

type DatabaseManager struct {
	session  *mgo.Session
	database *mgo.Database
}

func ConnectToDatabase(connectionString, databaseName string) (*DatabaseManager, error) {
	session, err := mgo.Dial(connectionString)

	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	dM := DatabaseManager{
		session,
		session.DB(databaseName),
	}

	return &dM, nil
}

func (s *DatabaseManager) InsertCertificate(cert *Certificate) error {
	err := s.database.C(certificatesCollection).Insert(cert)

	return err
}

func (s *DatabaseManager) InsertServerCertificate(cert *Certificate, serverName string) error {

	c, err := s.GetCertificateByCnameAndServer(cert.Cname, serverName)

	if err == nil {
		if c.Cname == cert.Cname {
			return errors.New("Cert already exists")
		}
	}

	server, err := s.GetServerByName(serverName)

	if err != nil {
		serv := Server{
			Name: serverName,
		}
		err = s.InsertServer(&serv)

		if err != nil {
			//error inserting new server
		}
		server, err = s.GetServerByName(serverName)

		if err != nil {
			return err
		}

	}

	cert.ServerId = server.ID
	cert.ServerName = server.Name

	err = s.InsertCertificate(cert)

	return err

}

func (s *DatabaseManager) GetCertificateById(id bson.ObjectId) (*Certificate, error) {
	result := Certificate{}
	err := s.database.C(certificatesCollection).FindId(id).One(&result)

	return &result, err
}

func (s *DatabaseManager) GetCertificates() ([]Certificate, error) {
	var result []Certificate
	err := s.database.C(certificatesCollection).Find(nil).All(&result)

	return result, err
}

func (s *DatabaseManager) DeleteCertificateById(id bson.ObjectId) error {
	err := s.database.C(certificatesCollection).RemoveId(id)

	return err
}

func (s *DatabaseManager) InsertServer(serv *Server) error {
	err := s.database.C(ServersCollection).Insert(serv)

	return err
}

func (s *DatabaseManager) GetServerById(id bson.ObjectId) (*Server, error) {
	result := Server{}
	err := s.database.C(ServersCollection).FindId(id).One(&result)

	return &result, err
}

func (s *DatabaseManager) GetServerByName(name string) (*Server, error) {
	result := Server{}
	err := s.database.C(ServersCollection).Find(bson.M{"name": name}).One(&result)

	return &result, err
}

func (s *DatabaseManager) GetCertificateByCnameAndServer(cname string, servername string) (*Certificate, error) {

	serv, err := s.GetServerByName(servername)

	if err != nil {
		return nil, err
	}

	result := Certificate{}
	err = s.database.C(certificatesCollection).Find(bson.M{"cname": cname, "serverid": serv.ID}).One(&result)

	return &result, err
}

func (s *DatabaseManager) GetServers() ([]Server, error) {
	var result []Server
	err := s.database.C(ServersCollection).Find(nil).All(&result)

	return result, err
}

func (s *DatabaseManager) DeleteServerById(id bson.ObjectId) error {
	err := s.database.C(ServersCollection).RemoveId(id)

	return err
}

func (s *DatabaseManager) GetServerCertificates(id bson.ObjectId) ([]Certificate, error) {
	var result []Certificate
	err := s.database.C(certificatesCollection).Find(bson.M{"serverid": id}).All(&result)
	return result, err
}

func (s *DatabaseManager) Close() {
	s.session.Close()
}
