package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var httpResponsWriter *http.Request
var tweetPrefix = "/and/"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
		httpResponsWriter = r
		parseTweet(html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseTweet(s string) {
	var output string
	if strings.HasPrefix(s, tweetPrefix) {
		fmt.Println(s)

		runes := []rune(s)

		for i := 0; i < len(runes); i++ {
			if i > 4 {
				output = output + string(runes[i])
			}
		}

		fmt.Println(strings.Join(strings.Split(output, "+"), " "))
		Tweet(strings.Join(strings.Split(output, "+"), " "))
	}

}

//Tweet Creates a new tweet
func Tweet(message string) {

	/* Insert API keys here */
	var consumerKey = ""
	var consumerSecret = ""
	var accessToken = ""
	var accessSecret = ""

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		println("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	//Update (POST!) Tweet (uncomment to run)
	var tweetURL, _, _ = client.Statuses.Update(message, nil)
	println("https://twitter.com/androidFinal/Status/" + tweetURL.IDStr)

}
