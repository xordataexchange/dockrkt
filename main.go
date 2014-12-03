package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type container struct {
	repository string
	user       string
	image      string
	tag        string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{repository}/{user}/{image}:{tag}", RocketHandler)
	http.Handle("/", r)
}
func RocketHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	repo, ok := vars["repository"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	usr, ok := vars["user"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	img, ok := vars["image"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
	}
	t, ok := vars["tag"]

	c := container{
		repository: repo,
		user:       usr,
		image:      img,
		tag:        t,
	}

}

func pullImage() {}

func runImage() {}

func inspectImage() {}

func exportImage() {}

func rmImage() {}

func createManifest() {}

func packageAci() {}
