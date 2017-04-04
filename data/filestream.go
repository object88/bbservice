package data

import (
	"bytes"
	"net/http"
)

func StreamImage(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("WHATEVER")
	w.Write(buffer.Bytes())
}
