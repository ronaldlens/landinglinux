package main

import (
	"encoding/json"
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

	err = json.Unmarshal([]byte(jsonString), &cards)
	if err != nil {
		panic(err)
	}
	logger.Println("Loaded data")
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	readData()
	err := tpl.Execute(w, cards)
	if err != nil {
		panic(err)
	}

	logger.Print("Render index")
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

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}

}

func main() {
	readData()
	startServer()
}
