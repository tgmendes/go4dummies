package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Film struct{}
type Species struct{}
type Vehicle struct{}
type Starship struct{}

// START OMIT
// People represents each individual from the API response.
type People struct {
	Name      string     `json:"name"`
	Height    string     `json:"height"`
	Mass      string     `json:"mass"`
	HairColor string     `json:"hair_color"`
	SkinColor string     `json:"skin_color"`
	EyeColor  string     `json:"eye_color"`
	BirthYear string     `json:"birth_year"`
	Gender    string     `json:"gender"`
	Homeworld string     `json:"homeworld"`
	Films     []Film     `json:"films"`     // HL
	Species   []Species  `json:"species"`   // HL
	Vehicles  []Vehicle  `json:"vehicles"`  // HL
	Starships []Starship `json:"starships"` // HL
	Created   time.Time  `json:"created"`
	Edited    time.Time  `json:"edited"`
	URL       string     `json:"url"`
}

// PeopleAPI represents the API response for the people.
type PeopleAPI struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []People `json:"results"` // HL
}

// END OMIT

func main() {
	pr := &PeopleAPI{}

	b, _ := ioutil.ReadFile("sw_sample.json")

	err := json.Unmarshal(b, pr) // HL
	if err != nil {
		log.Fatal("unable to unmarshal ", err.Error())
	}

	for _, p := range pr.Results {
		fmt.Println(p.Name)
	}
}
