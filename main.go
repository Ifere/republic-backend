package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"republic-backend/config/db"
	"republic-backend/config/middleware"
)

// each start method requires the instance of the app handler
//TODO refactor register service to accept various db connection types
func registerServices(router *mux.Router, conn db.MongoDB) {

	InitializeUser(conn).Start(router)

	InitializeZone(conn).Start(router)

	InitializeMember(conn).Start(router)


}


func registerMiddleWares(router *mux.Router) *negroni.Negroni {
	logger()
	n := negroni.Classic()
	n.Use(middleware.Cors())
	n.UseHandler(router)
	return n
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	var db db.MongoDB
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	_ = registerMiddleWares(router)
	registerServices(router, db)

	PORT := "8080"

	log.Printf("server running on port %v", port)

	err = http.ListenAndServe(":"+PORT, router)

	fmt.Println(err)

}

func logger() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.SetOutput(logger.Writer())
}