package main

import (
	"log"
	"os"

	"github.com/levigross/grequests"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{
	Auth: []string{GITHUB_TOKEN, "x-oauth-basic"},
	Headers: map[string]string{
		"Accept": "application/vnd.github.v3+json",
	},
}

type Repo struct {
	ID       int    `json: "id"`
	Name     string `json: "name"`
	FullName string `json: "full_name"`
	Forks    int    `json: "forks"`
	Private  bool   `json: "private"`
}

func GetStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, requestOptions)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp
}

func main() {
	var repos []Repo
	var repoUrl = "https://api.github.com/users/ddefrancesco/repos?per_page=50"
	resp := GetStats(repoUrl)
	resp.JSON(&repos)
	var count = 0
	var countPrivate = 0
	for _, v := range repos {
		if v.Private {
			countPrivate++
		}
		count++
		log.Printf("Nome Repository: %s", v.Name)
	}
	log.Printf("Trovati %d repository, di cui %d privati", count, countPrivate)

}
