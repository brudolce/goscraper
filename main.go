package main

import (
	"bufio"
	"fmt"
	"log"
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
	task1(doc)

	//TASK 2 - Find Page title
	task2(doc)

	//TASK 3 - Headings count by level
	task3(doc)

	//TASK 4 - Amount of internal and external links
	//TASK 5 - Amount of inacessible links
	task4n5(doc)

	//TASK 6 - If a Page contains a login form
	task6(doc)

	fmt.Println()

}
