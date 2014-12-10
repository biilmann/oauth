// This is a runnable example of making an oauth authorized request
// to the Twitter api
// Enter your consumer key/secret and token key/secret as command line flags
// ex: go run example.go -ck ABC -cs DEF -tk 123 -ts 456

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nhjk/oauth"
)

var ck = flag.String("ck", "", "consumer key")
var cs = flag.String("cs", "", "consumer secret")
var tk = flag.String("tk", "", "token key")
var ts = flag.String("ts", "", "token secret")

func main() {
	flag.Parse()

	// create an http client and a request for it to send
	client := new(http.Client)
	req, _ := http.NewRequest("GET", "https://api.twitter.com/1.1/statuses/home_timeline.json", nil)

	// a consumer allows you to authorize requests with a token
	cons := oauth.Consumer{*ck, *cs}

	// authorize request
	cons.Authorize(req, &oauth.Token{*tk, *ts})

	// send request and print body
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", body)
}
