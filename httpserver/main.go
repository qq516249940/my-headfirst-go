package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	os.Mkdir("file", 0777)

	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("file"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
