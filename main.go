package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("Error: Not Found"))
}

func InitRouter() *httprouter.Router{
	router := httprouter.New()

	webDir := os.Getenv("WEB_DIR")
	if webDir == ""{
		webDir = "./"
	}


	/* Static file router */
	fileServer := http.FileServer(http.Dir("./"))
	router.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		r.URL.Path = p.ByName("filepath")
		fileServer.ServeHTTP(w, r)
	})

	router.NotFound = http.HandlerFunc(NotFoundHandler)
	return router
}

func main() {
	bindPort := ":80"
	bindEnv := os.Getenv("BIND_PORT")
	if bindEnv != ""{
		bindPort = fmt.Sprintf(":%s",bindEnv)
	}
	log.Fatal(http.ListenAndServe(bindPort, InitRouter()))
}
