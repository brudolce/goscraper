# Go Scraper

![](https://img.shields.io/badge/go-v1.14-blue) ![](https://img.shields.io/badge/goquery-v1.5.1-blue) ![](https://img.shields.io/badge/Dockerfile-golang:alpine-green)

This is a simple scraper to check information from webpages.

[TOCM]

### Features

- Checks HTML version
- Checks page title
- Headings count by level
- Amount of internal and external Links
- Amount of inacessible links
- If a page contains a login form

---

### Get Started

##### Docker

If you have docker installed, clone this repository anywhere in your machine and in terminal, on this project folder, type the following comands:

```bash
 docker build -t goscraper .
 docker run -it goscraper
```

##### Golang

---

If you do not have docker installed but have golang in your machine, clone this repo inside your \$GOPATH work directory, them run the application in the terminal on this project folder:

```bash
go run .
```

---

### Extras

install docker: https://docs.docker.com/get-docker/
install golang: https://golang.org/doc/install/
