package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teksoftgroup/embed-solidjs/client"
)

// serve the spa without embedding it.
// func main() {
// 	http.HandleFunc("/", handleSpa)

// 	log.Println("Server is listening to post 5050")
// 	log.Fatal(http.ListenAndServe(":5050", nil))
// }

// func handleSpa(w http.ResponseWriter, r *http.Request) {
// 	rootPath, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	buildPath := filepath.Join("client", "dist")
// 	path := filepath.Join(rootPath, buildPath, r.URL.Path)

// 	_, err = os.Stat(path)
// 	if os.IsNotExist(err) {
// 		http.ServeFile(w, r, filepath.Join(buildPath, "index.html"))
// 		return
// 	} else if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	http.FileServer(http.Dir(buildPath)).ServeHTTP(w, r)
// }

// serve the spa in embedded binary
// let go know where to get all the files to embed
// using the embed directive
//

func main() {
	http.HandleFunc("/", handleSpa)

	log.Printf("Server is listening to post %s", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil))
}

func handleSpa(w http.ResponseWriter, r *http.Request) {
	buildPath := "dist"

	path, err := client.BuildFS.Open(filepath.Join(buildPath, r.URL.Path))
	if os.IsNotExist(err) {
		index, err := client.BuildFS.ReadFile(filepath.Join(buildPath, "index.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer path.Close()

	http.FileServer(client.BuildHTTPFS()).ServeHTTP(w, r)
}
