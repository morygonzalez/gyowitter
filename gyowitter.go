package main

import (
	"flag"
	"log"
	"net/url"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kentaro/gyo"
)

var apiToken = flag.String("token", "", "Yo API Token")
var port = flag.Int("port", 8080, "Port of the server")
var path = flag.String("path", "/callback", "Callback URL of the server")

func init() {
	flag.Parse()
}

func main() {
	if *apiToken == "" {
		log.Fatalln("API token is required")
	}

	gyo := gyo.NewGyo(*apiToken)
	gyo.Server(*path, *port, func(username string) {
		gyo.Yo(username)
		log.Printf("Sent Yo to %s\n", username)

		postToTwitter(username)
	})

	return
}

func postToTwitter(username string) {
	anaconda.SetConsumerKey("consumer-key")
	anaconda.SetConsumerSecret("consumer-secret")
	api := anaconda.NewTwitterApi("secret-key", "secret-token")

	v := url.Values{}
	strToPost := "また @" + strings.ToLower(username) + " さんから Yo がありました。仕事してるんですかね？"

	api.PostTweet(strToPost, v)

	return
}
