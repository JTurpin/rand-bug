package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	players [4]string
)

type Game struct {
	s1b string
	s1w string
	s2b string
	s2w string
}

func main() {

	fmt.Println("running on: http://localhost:9090 for testing")
	http.HandleFunc("/", index)              // setting router rule
	http.HandleFunc("/style.css", style)     // setting router rule for style
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// handler for requests, GET and POST
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html.template")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		players[0] = r.Form["user1"][0]
		players[1] = r.Form["user2"][0]
		players[2] = r.Form["user3"][0]
		players[3] = r.Form["user4"][0]

		fmt.Println("Ordered: ", players)

		t := time.Now()
		rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

		for i := len(players) - 1; i > 0; i-- {
			j := rand.Intn(i)
			players[i], players[j] = players[j], players[i]
		}
		fmt.Println("Shuffled: ", players)

		fmt.Fprintf(w, "Bathroom White: ")
		fmt.Fprintln(w, players[0])
		fmt.Fprintf(w, "Bathroom Black: ")
		fmt.Fprintln(w, players[1])
		fmt.Fprintf(w, "Conf room White: ")
		fmt.Fprintln(w, players[2])
		fmt.Fprintf(w, "Conf room Black: ")
		fmt.Fprintln(w, players[3])

	}
}

// serve: code.css
func style(w http.ResponseWriter, r *http.Request) {
	css, err := ioutil.ReadFile("style.css")
	check(err)
	fmt.Fprintln(w, string(css))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
