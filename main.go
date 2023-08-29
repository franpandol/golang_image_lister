package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ImageList struct {
	Images []string `json:"images"`
}

func main() {
    router := mux.NewRouter()

    // Add your API routes here
	router.HandleFunc("/api/images/{filename}", imageHandler).Methods("GET")
	router.HandleFunc("/api/images", imageListHandler).Methods("GET")
    // Enable CORS
    corsMiddleware := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Allow requests from any origin
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    })
    router.Use(corsMiddleware.Handler)

    // Start the server
    http.ListenAndServe(":8086", router)
 

}

func imageHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    filename := vars["filename"]
	fmt.Println(filename)
    http.ServeFile(w, r, "./images/" + filename)
}

func imageListHandler(w http.ResponseWriter, r *http.Request) {
    // Set the appropriate content type header
    w.Header().Set("Content-Type", "application/json")

    // Get the data
    data := getData()

    // Write the data to the response
    json.NewEncoder(w).Encode(data)
}
func getData() ImageList {
	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	var imgs []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".jpg") || strings.HasSuffix(f.Name(), ".png") {
			imgs = append(imgs, f.Name())
		}
	}

	il := ImageList{Images: imgs}
	fmt.Println(il)
	return il
}
