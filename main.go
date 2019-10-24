package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/sachin-mundhe/recover-middleware/middleware"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/debug/", Debug)
	router.HandleFunc("/panic/", panicDemo)
	router.HandleFunc("/panic-after/", panicAfterDemo)
	log.Fatalln(http.ListenAndServe(":3000", middleware.Recover(router, true)))

}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")

}

//Debug This method handles debug route
func Debug(w http.ResponseWriter, r *http.Request) {
	path := "/home/gslab/Desktop/RecoverMiddleware/main.go"
	file, err := os.Open(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = quick.Highlight(w, b.String(), "go", "html", "monokai")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
