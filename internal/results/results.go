package results

import (
	"log"
	"sync"

	"github.com/mmcdole/gofeed"
)

func WatchRss(ch chan MyRSS, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	res := MyRSS{}
	rssEntry := Entry{}

	newParser := gofeed.NewParser()
	newFeed, err := newParser.ParseURL(url)
	if err != nil {
		log.Printf("Erreur pour récup le RSS de : %s\n%s", url, err)
		res.Website = url
		rssEntry.Published = "0"
		rssEntry.Title = "Impossible d'avoir le flux RSS"
		rssEntry.Url = "0"
		res.Entries = append(res.Entries, rssEntry)
	} else {
		res.Website = newFeed.Title
		for _, item := range newFeed.Items {
			rssEntry.Published = item.Published
			rssEntry.Title = item.Title
			rssEntry.Url = item.Link
			res.Entries = append(res.Entries, rssEntry)
		}
	}

	ch <- res
}
