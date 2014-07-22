package main

import (
	"bufio"
	"flag"

	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"

	"time"

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
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")
	api := anaconda.NewTwitterApi("", "")

	v := url.Values{}
	suffix := getSuffix()
	strToPost := "また @" + strings.ToLower(username) + " さんから Yo がありました。" + suffix
	log.Printf(strToPost)

	api.PostTweet(strToPost, v)

	return
}

func getSuffix() string {
	rand.Seed(time.Now().UTC().UnixNano())

	// open input file
	fi, err := os.Open("suffixes.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read scanner
	scanner := bufio.NewScanner(fi)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines[rand.Intn(len(lines))]
}
