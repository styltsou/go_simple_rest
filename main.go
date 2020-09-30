package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Articles struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}

// Save articles slice to json file

// Return Articles struct containing the json file's data

// Get all articles
func getArticles(writer http.ResponseWriter, req *http.Request) {
	// Open articles.json
	articlesJson, err := os.Open("articles.json")
	// log error if there is one
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of the json file so that can be parsed later
	defer articlesJson.Close()

	var articles Articles
	byteArray, _ := ioutil.ReadAll(articlesJson)
	json.Unmarshal(byteArray, &articles)

	json.NewEncoder(writer).Encode(articles)
}

// Get an article based on its id
func getArticle(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	// Open articles.json
	articlesJson, err := os.Open("articles.json")
	// log error if there is one
	if err != nil {
		fmt.Println(err)
	}

	defer articlesJson.Close()

	var articles Articles
	byteArray, _ := ioutil.ReadAll(articlesJson)
	json.Unmarshal(byteArray, &articles)

	for _, article := range articles.Articles {
		if article.Id == id {
			json.NewEncoder(writer).Encode(article)
			return
		}
	}

	fmt.Println("Article not found")
}

// Add a new article to the json file
func addArticle(writer http.ResponseWriter, req *http.Request) {
	// Parse request body
	reqBody, _ := ioutil.ReadAll(req.Body)

	// Create a new article from request body

	var article Article
	json.Unmarshal(reqBody, &article)
}

func updateArticle(writer http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	// id := params["id"]

	// reqBody, _ := ioutil.ReadAll(req.Body)

	// var updatedArticle Article
	// json.Unmarshal(reqBody, &updatedArticle)
	// updatedArticle.Id = id

	// for idx, article := range Articles {
	// 	if article.Id == id {
	// 		Articles = append(Articles[:idx], Articles[idx+1:]...)
	// 		Articles = append(Articles, updatedArticle)
	// 	}
	// }

}

func deleteArticle(writer http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	// id := params["id"]

	// for idx, article := range Articles {
	// 	if article.Id == id {
	// 		Articles = append(Articles[:idx], Articles[idx+1:]...)
	// 	}
	// }
}

func handleRequests() {
	// Create a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	// Create endpoints
	router.HandleFunc("/articles", addArticle).Methods("POST")
	router.HandleFunc("/articles", getArticles)
	router.HandleFunc("/articles{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", getArticle)

	// Listen for requests
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	fmt.Println("Server listening on port 5000")

	// Create a global slice to store the articles
	// var articles Articles
	// Populate articles slice with json data

	handleRequests()
}
