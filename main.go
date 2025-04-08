package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heshanu/gorest/models"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {

	articles := models.Articles{
		models.Article{
			Title:   "Test",
			Content: ":Helo",
			Desc:    "desc",
		},
	}
	//fmt.Fprintf(w, "All ariticle endpoint hit")
	json.NewEncoder(w).Encode(articles)
}

func SaveArticle(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomePage)
	router.HandleFunc("articles", GetAllArticles).Methods("GET")
	log.Fatal((http.ListenAndServe(":8081", router)))
}

func main() {
	handleRequests()
	http.ListenAndServe(":8081", nil)
}
