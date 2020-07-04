package main

import (
	"math/rand"
	"time"
)

const (
	urlTopStories              = "https://api.nytimes.com/svc/topstories/v2/home.json"
	urlMostPopular             = "https://api.nytimes.com/svc/mostpopular/v2/viewed/1.json"
	urlHardcoverFictionReviews = "https://api.nytimes.com/svc/books/v3/lists/current/hardcover-fiction.json"
)

var urls = [...]string{urlTopStories,
	urlMostPopular,
	urlHardcoverFictionReviews,
}

func fetch(url string, channel chan<- string) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	duration := time.Duration(random.Intn(150)) * time.Millisecond
	time.Sleep(duration)
	channel <- url
}

func main() {

}
