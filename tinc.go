package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var client *http.Client
var increment int
var numreq int
var URLstring string

// main function
func main() {

	log.SetFlags(0) //supress date and time on each line
	flag.StringVar(&URLstring, "u", "", "--> URL with 'ZCZC' where the interger to increment should be placed "+"\n")
	flag.IntVar(&increment, "i", 0, "--> Interger to increment by in each request "+"\n")
	flag.IntVar(&numreq, "n", 0, "--> Number of requests to send "+"\n")
	//flag.PrintDefaults()
	flag.Parse()

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 30 * time.Second,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for i := 0; i < numreq; i++ {
		s := strconv.Itoa(increment)
		newURL := strings.Replace(URLstring, "ZCZC", s, 1)
		a, err := url.Parse(newURL)
		if err != nil {
		        fmt.Println(color.RedString(err.Error()))
		}

		fetchURL(a)
		increment += increment

	}
} //end main

func fetchURL(u *url.URL) {
	time_start := time.Now()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println(color.RedString(err.Error()))
	}
	req.Header.Set("User-Agent", "inc/0.1")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(color.RedString(err.Error()))
	}

	defer resp.Body.Close()
	log.Printf("%s: %v", color.GreenString(time.Since(time_start).String()), color.BlueString(u.String())+"\n\n")
	io.Copy(ioutil.Discard, resp.Body)

}
