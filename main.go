package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

type container struct {
	registry string
	image    string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{registry}/{image}", RocketHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
func RocketHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	reg, ok := vars["registry"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	img, ok := vars["image"]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
	}
	c := container{
		registry: repo,
		image:    img,
	}
	err := pullImage(c)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	return
}

func pullImage(c container) error {
	cmd := exec.Command("docker", "pull", nameFromC(c))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	return nil
}

func runImage() {}

func inspectImage() {}

func exportImage() {}

func rmImage() {}

func createManifest() {}

func packageAci() {}

func nameFromC(c container) string {
	if c.registry != nil {
		return fmt.Sprintf("%s/%s", c.registry, c.image)
	}
	return fmt.Sprintf("%s", c.image)
}
