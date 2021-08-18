package main

import (
	"fmt"
	"html/template"
	"log"
	myrss "mesRSS/internal/results"
	"net/http"
	"sync"
)

const Port = 8000

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	//urls := [1]string{"https://scantrad.net/rss/"}
	urls := [2]string{"https://scantrad.net/rss/", "https://www.reddit.com/r/OnePunchMan/.rss"}

	c := make(chan myrss.MyRSS, len(urls))
	defer close(c)

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go myrss.WatchRss(c, url, &wg)
	}

	wg.Wait()

	t, err := template.ParseFiles("../static/index.html")
	if err != nil {
		fmt.Printf("Can't open template : %s\n", err)
		return
	}

	tmp := make([]interface{}, 0)

	for len(c) > 0 {
		tmp = append(tmp, <-c)
	}

	t.Execute(w, tmp)
}

func main() {
	fmt.Printf("Démarrage sur le port %d\n", Port)

	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)

	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}
}
