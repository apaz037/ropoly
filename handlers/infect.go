package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

func InfectHandler(w http.ResponseWriter, r *http.Request) {
	f1, _ := ioutil.ReadDir("/tmp")
	d1 := []byte("foo")
	err := ioutil.WriteFile(fmt.Sprintf("/tmp/data%03d", len(f1) + 1), d1, 0644)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	files, _ := ioutil.ReadDir("/tmp")

	io.WriteString(w, fmt.Sprintf("%v\n", len(files)))
}
