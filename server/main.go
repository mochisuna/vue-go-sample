package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/unrolled/render"
)

type MonsterID int
type MonsterHP int
type MonsterName string

type Monster struct {
	ID   MonsterID   `json:"id"`
	Name MonsterName `json:"name"`
	Max  MonsterHP   `json:"max"`
	HP   MonsterHP   `json:"hp"`
}

func main() {
	rendering := render.New(render.Options{})
	raw, err := ioutil.ReadFile("./list.json")
	if err != nil {
		log.Printf("Failed ReadFile. err: %v", err)
		os.Exit(1)
	}

	mons := []Monster{}
	if err := json.Unmarshal(raw, &mons); err != nil {
		log.Printf("Failed Unmarshal. err: %v", err)
		os.Exit(1)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/monsters", func(w http.ResponseWriter, r *http.Request) {
		for _, mon := range mons {
			log.Printf("%v: %v (%v/%v)\n", mon.ID, mon.Name, mon.HP, mon.Max)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		rendering.JSON(w, http.StatusOK, mons)
	})
	log.Println("Start Server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Failed ListenAndServe. err: %v", err)
		os.Exit(1)
	}
}
