package main

import (
	"os"
	"io/ioutil"
	"strings"
	"time"
)

type Identifier struct {
	threads int // amount of simultaneously running http requests
	urlChan chan string // channel where we send urls to visit
	timeout int // timeout for http requests
	platforms []Platform // footprints name - footprints
	visited *os.File // visitedPath
	errors *os.File // errorPath
}

func NewIdentifier(threads, t int) Identifier {
	vF, eF := GetVisitAndErrorFiles()
	return Identifier{
		threads: threads,
		urlChan: make(chan string, 10000),
		timeout: t,
		platforms: ParseFootprints(),
		visited: vF,
		errors: eF,
	}
}

func (pi *Identifier) Start() {
	for i := 0; i < pi.threads; i++ {
		visitor := Visitor{}
		visitor.Start(pi)
	}
}

func main() {
	iden := NewIdentifier(20, 30)
	urls, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	urlsSlice := strings.Split(string(urls), "\n")
	iden.Start()
	for _, u := range urlsSlice {
		iden.urlChan <- strings.TrimSpace(u)
	}
	time.Sleep(time.Hour * 333)
}