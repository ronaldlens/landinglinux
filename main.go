package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type CardsData struct {
	Cards []struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Text     string `json:"text"`
		Button   string `json:"button"`
		URL      string `json:"url"`
	} `json:"cards"`
}

var cards CardsData
var logger = log.New(os.Stderr, "http: ", log.LstdFlags)

func readData() {
	dat, err := os.ReadFile("data/cards.json")
	if err != nil {
		panic(err)
	}
	jsonString := string(dat)
	cards = CardsData{}

	json.Unmarshal([]byte(jsonString), &cards)
	logger.Println("Loaded data")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	readData()
	tpl.Execute(w, cards)
	logger.Print("Render index")
}

func Sample() string {
	fmt.Println("test from web")
	return "hello"
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	mux := http.NewServeMux()
	fsData := http.FileServer(http.Dir("data"))
	mux.Handle("/data/", http.StripPrefix("/assets/", fsData))
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)
}

func main() {
	readData()
	startServer()
}
