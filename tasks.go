package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

//TASK 1 - HTML Version ??
func task1(doc *goquery.Document, wg *sync.WaitGroup) {
	charset := doc.Find("charset").Contents().Text()
	fmt.Println("\n\n HTML Version: 5 ", charset)
	wg.Done()
}

//TASK 2 - Find Page title
func task2(doc *goquery.Document, wg *sync.WaitGroup) {
	var pageTitle string = doc.Find("title").Contents().Text()
	fmt.Println("\n PageTitle :", pageTitle)
	wg.Done()
}

//TASK 3 - Headings count by level
func task3(doc *goquery.Document, wg *sync.WaitGroup) {
	var h1Count, h2Count, h3Count, h4Count, h5Count, h6Count = 0, 0, 0, 0, 0, 0

	doc.Find("h1").Each(func(index int, element *goquery.Selection) { h1Count++ })
	doc.Find("h2").Each(func(index int, element *goquery.Selection) { h2Count++ })
	doc.Find("h3").Each(func(index int, element *goquery.Selection) { h3Count++ })
	doc.Find("h4").Each(func(index int, element *goquery.Selection) { h4Count++ })
	doc.Find("h5").Each(func(index int, element *goquery.Selection) { h5Count++ })
	doc.Find("h6").Each(func(index int, element *goquery.Selection) { h6Count++ })
	fmt.Println("\n Headings count by level:\n    H1 Tags: ", h1Count, "\n    H2 Tags: ", h2Count, "\n    H3 Tags: ", h3Count, "\n    H4 Tags: ", h4Count, "\n    H5 Tags: ", h5Count, "\n    H6 Tags: ", h6Count)
	wg.Done()
}

//TASK 4 - Amount of internal and external links
func task4(doc *goquery.Document, internal []string, external []string, wg *sync.WaitGroup) {
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
	go task5(external, wg)
}

//TASK 5 - Amount of inacessible links
func task5(links []string, wg *sync.WaitGroup) {
	var inacessible int = 0
	fmt.Println("\n        ...checking links ... ")

	for _, link := range links {
		_, err := http.Get(link)
		if err != nil {
			inacessible++
			fmt.Println("    Invalid link : ", link)
		}
	}
	fmt.Println(" Amount of inacessible links: ", inacessible)
	wg.Done()
}

//TASK 6
func task6(doc *goquery.Document, wg *sync.WaitGroup) {
	var loginForm bool = false
	doc.Find("input").Each(func(index int, element *goquery.Selection) {
		pass, _ := element.Attr("type")
		if pass == "password" {
			loginForm = true
		}
	})
	fmt.Println("\n Page contains a login form: ", loginForm)
	wg.Done()
}
