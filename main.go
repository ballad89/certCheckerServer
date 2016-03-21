package main

import (
	"certcheckerServer/handlers"
	"certcheckerServer/hub"
	"certcheckerServer/libhelper"
	"certcheckerServer/models"
	"github.com/desertbit/glue"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	connectionString := libhelper.GetEnvOrDefault("CONNECTIONSTRING", "mongodb://localhost")

	databaseName := libhelper.GetEnvOrDefault("DATABASENAME", "certchecker")

	redisUrl := libhelper.GetEnvOrDefault("REDISURL", "redis://localhost")

	db, err := models.ConnectToDatabase(connectionString, databaseName)

	if err != nil {
		panic(err)
	}

	env := &handlers.Env{
		Db: db,
	}

	defer env.Db.Close()

	glueServer := glue.NewServer(glue.Options{
		HTTPSocketType: glue.HTTPSocketTypeNone,
		HTTPHandleURL:  "/ws/",
	})

	err = hub.InitHub(redisUrl)

	if err != nil {
		panic(err)
	}

	glueServer.OnNewSocket(hub.HandleSocket)

	r := mux.NewRouter()

	r.HandleFunc("/certificates", env.GetCerticates).Methods("GET")
	r.HandleFunc("/certificate/{id}", env.GetCertificate).Methods("GET")
	r.HandleFunc("/certificates", env.PostCertificate).Methods("POST")

	r.HandleFunc("/servers", env.GetServers).Methods("GET")
	r.HandleFunc("/server/{id}", env.GetServer).Methods("GET")
	r.HandleFunc("/server/certificates/{id}", env.GetServerCertificates).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)
	http.Handle("/ws/", glueServer)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	http.ListenAndServe(":3000", nil)
}
