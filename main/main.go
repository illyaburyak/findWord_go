//package main
//
//import (
//	"fmt"
//	"github.com/PuerkitoBio/goquery"
//
//	"net/http"
//	"strings"
//)
//
//type findWord int
//
//// This func is going to write back respond
//func (m findWord) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//
//	fmt.Println("Enter Word you wanna search for ")
//	var userWord string
//	fmt.Scanln(&userWord)
//
//	url := req.URL
//
//	newUrl := url.String()
//	//get html
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(newUrl))
//	if err != nil {
//		panic(err)
//	}
//
//	// find user name
//	doc.Find(userWord).Each(func(_ int, node *goquery.Selection) {
//		fmt.Print("Your word is " + userWord)
//	})
//
//}
//
//
//func main() {
//	// display output in the next line
//	fmt.Println("Enter URL please: ")
//	var userUrl string
//
//	// Storing input from user
//	fmt.Scanln(&userUrl)
//
//
//	var searchingFunc findWord
//	err := http.ListenAndServe(userUrl, searchingFunc)
//	if err != nil {
//		return
//	}
//
//}

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	// get user url
	fmt.Println("Enter URL please: ")
	var userUrl string

	// get user word
	fmt.Println("Enter you word please: ")
	var userWord string

	// trim spaces
	userWord = strings.TrimSpace(userWord)

	// fetch data
	resp, err := http.Get(userUrl)

	if err != nil {
		log.Fatal(err)
	}

	// close body when finish
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// read content
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Fatal(err)
	}

	// storing html
	htmlPage := string(body)

	// get dom
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlPage))
	if err != nil {
		panic(err)
	}

	// find user word
	doc.Find(userWord).Each(func(_ int, node *goquery.Selection) {
		fmt.Print("Your word is " + userWord)
		ret := node.Find(userWord)
		fmt.Println(ret)
	})

}
