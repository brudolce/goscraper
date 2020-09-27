package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	//Get from terminal the site url
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter site url to be scraped : \n\n   do not forget the 'https://', e.g. https://www.home24.com/ or https://google.com\n\n  --> ")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	//SETUP
	doc, err := goquery.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	//TASK 1 - HTML Version ??
	charset := doc.Find("charset").Contents().Text()
	fmt.Println("\n\n HTML Version: 5 ", charset)

	//TASK 2 - Find Page title
	var pageTitle string = doc.Find("title").Contents().Text()
	fmt.Println("\n PageTitle :", pageTitle)

	//TASK 3 - Headings count by level
	var h1Count, h2Count, h3Count, h4Count, h5Count, h6Count = 0, 0, 0, 0, 0, 0

	doc.Find("h1").Each(func(index int, element *goquery.Selection) { h1Count++ })
	doc.Find("h2").Each(func(index int, element *goquery.Selection) { h2Count++ })
	doc.Find("h3").Each(func(index int, element *goquery.Selection) { h3Count++ })
	doc.Find("h4").Each(func(index int, element *goquery.Selection) { h4Count++ })
	doc.Find("h5").Each(func(index int, element *goquery.Selection) { h5Count++ })
	doc.Find("h6").Each(func(index int, element *goquery.Selection) { h6Count++ })
	fmt.Println("\n Headings count by level:\n    H1 Tags: ", h1Count, "\n    H2 Tags: ", h2Count, "\n    H3 Tags: ", h3Count, "\n    H4 Tags: ", h4Count, "\n    H5 Tags: ", h5Count, "\n    H6 Tags: ", h6Count)

	//TASK 3 - Amount of internal and external links
	var internal, external = []string{}, []string{}

	doc.Find("body a").Each(func(index int, element *goquery.Selection) {
		link, _ := element.Attr("href")
		if strings.Contains(link, "http") {
			external = append(external, link)
		} else {
			internal = append(internal, link)
		}
		internal = uniqueStringArray(internal)
		external = uniqueStringArray(external)
	})
	fmt.Println("\n Amount of External Links: ", len(external), "\n Amount of Internal Links: ", len(internal))

	//TASK4 - Amount of inacessible links
	var inacessible int = 0
	fmt.Println("\n        ...checking links ... ")

	for _, link := range external {
		_, err := http.Get(link)
		if err != nil {
			inacessible++
			fmt.Println("    Invalid link : ", link)
		}
	}
	fmt.Println(" Amount of inacessible links: ", inacessible)

	//TASK5 - If a Page contains a login form
	var loginForm bool = false
	doc.Find("input").Each(func(index int, element *goquery.Selection) {
		pass, _ := element.Attr("type")
		if pass == "password" {
			loginForm = true
		}
	})

	fmt.Println("\n Page contains a login form: ", loginForm)
}
