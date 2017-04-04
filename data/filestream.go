package data

import (
	"io/ioutil"
	"log"
	"net/http"
)

func StreamImage(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadFile("./DSC_5537.jpg")
	if err != nil {
		log.Fatalf("Got error '%s' while reading image file\n", err)
		panic(err)
	}
	w.Write(buf)
}
