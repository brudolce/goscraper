package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

//TASK 1 - HTML Version ??
func task1(doc *goquery.Document) {
	charset := doc.Find("charset").Contents().Text()
	fmt.Println("\n\n HTML Version: 5 ", charset)

}

//TASK 2 - Find Page title
func task2(doc *goquery.Document) {
	var pageTitle string = doc.Find("title").Contents().Text()
	fmt.Println("\n PageTitle :", pageTitle)

}

//TASK 3 - Headings count by level
func task3(doc *goquery.Document) {
	var h1Count, h2Count, h3Count, h4Count, h5Count, h6Count = 0, 0, 0, 0, 0, 0

	doc.Find("h1").Each(func(index int, element *goquery.Selection) { h1Count++ })
	doc.Find("h2").Each(func(index int, element *goquery.Selection) { h2Count++ })
	doc.Find("h3").Each(func(index int, element *goquery.Selection) { h3Count++ })
	doc.Find("h4").Each(func(index int, element *goquery.Selection) { h4Count++ })
	doc.Find("h5").Each(func(index int, element *goquery.Selection) { h5Count++ })
	doc.Find("h6").Each(func(index int, element *goquery.Selection) { h6Count++ })
	fmt.Println("\n Headings count by level:\n    H1 Tags: ", h1Count, "\n    H2 Tags: ", h2Count, "\n    H3 Tags: ", h3Count, "\n    H4 Tags: ", h4Count, "\n    H5 Tags: ", h5Count, "\n    H6 Tags: ", h6Count)

}

//TASK 4 - Amount of internal and external links
func task4n5(doc *goquery.Document) {
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
	go task5(external)
}

//TASK 5 - Amount of inacessible links
func urlCallCount(links []string) int {
	count := 0
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				count++
				fmt.Println("    Invalid link : ", url)
			}
		}(link)
	}
	wg.Wait()
	return count
}

func task5(links []string) {
	inacessible := urlCallCount(links)
	fmt.Println("\n        ...checking links ... ")
	fmt.Println(" Amount of inacessible links: ", inacessible)

}

//TASK 6
func task6(doc *goquery.Document) {
	var loginForm bool = false
	doc.Find("input").Each(func(index int, element *goquery.Selection) {
		pass, _ := element.Attr("type")
		if pass == "password" {
			loginForm = true
		}
	})
	fmt.Println("\n Page contains a login form: ", loginForm)

}
