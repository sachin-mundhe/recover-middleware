package main

import (
	"fmt"
	"log"
	"net/http"

	m "github.com/sachin-mundhe/recover-middleware/route-handler"

	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("Init from main")
}

var listenAndServeFunc = http.ListenAndServe

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", m.HomepageHandler)
	router.HandleFunc("/sourcecode/", m.SourceCodeHandler)
	router.HandleFunc("/panic/", m.PanicHandler)
	log.Fatalln(listenAndServeFunc(":3000", m.Recover(router, true)))
}

//localhost:3000/sourcecode/?line=ewr&path=/home/gslab/Desktop/InstError.txt
