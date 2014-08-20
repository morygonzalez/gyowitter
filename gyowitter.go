package main

import (
	"encoding/json"
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

var port = flag.Int("port", 8080, "Port of the server")
var path = flag.String("path", "/callback", "Callback URL of the server")
var config Config

type Config struct {
	API              map[string]map[string]string
	Suffixes         map[string][]string
	UsernameMappings map[string]string
}

func init() {
	flag.Parse()
	config = Load("config.json")
}

func main() {
	gyo := gyo.NewGyo(config.API["Yo"]["token"])
	gyo.Server(*path, *port, func(username string) {
		gyo.Yo(username)
		log.Printf("Sent Yo to %s\n", username)

		postToTwitter(username)
	})

	return
}

func Load(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("You need config.json to put on same directory.")
	}

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("You need valid config.json. Check the syntax.")
	}
	return config
}

func postToTwitter(username string) {
	configTwitter := config.API["Twitter"]
	anaconda.SetConsumerKey(configTwitter["APIKey"])
	anaconda.SetConsumerSecret(configTwitter["APISecret"])
	api := anaconda.NewTwitterApi(configTwitter["AccessToken"], configTwitter["AccessTokenSecret"])

	v := url.Values{}
	suffix := getSuffix()
	username = usernameMapping(strings.ToLower(username))
	strToPost := "また @" + username + " さんから Yo がありました。" + suffix
	log.Printf(strToPost)

	api.PostTweet(strToPost, v)

	return
}

func getSuffix() string {
	suffixes := config.Suffixes
	rand.Seed(time.Now().UTC().UnixNano())

	var suffixKind string
	if rand.Intn(len(suffixes)) == 0 {
		suffixKind = "Accuse"
	} else {
		suffixKind = "Poem"
	}

	return suffixes[suffixKind][rand.Intn(len(suffixes[suffixKind]))]
}

func usernameMapping(username string) string {
	mappings := config.UsernameMappings
	if mapped, ok := mappings[username]; ok {
		return mapped
	} else {
		return username
	}
}
