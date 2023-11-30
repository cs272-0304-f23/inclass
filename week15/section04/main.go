package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
)

type Book struct {
	Title, Author string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("ParseFiles: ", err)
	}

	books := []Book{
		{"The Iliad", "Homer"},
		{"Dracula", "Bram Stoker"},
	}

	err = t.Execute(w, books)
	if err != nil {
		log.Fatalln("Execute: ", err)
	}
}

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/", handleIndex)
	http.Handle("/project06.css", http.FileServer(http.Dir("./")))
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	<-exit
	log.Println("clean shutdown here")
}
