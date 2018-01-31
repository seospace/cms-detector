package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

// visitor makes http requests and identifies websites by footprints
type Visitor struct {
}

// starts the visitor loop
func (v *Visitor) Start(id *Identifier) {
	client := http.Client{Timeout: ConvertTimeout(id.timeout)}
	go func() {
		for {
			select {
			case work := <-id.urlChan:
				fmt.Println("Received request to visit:", work)
				fmt.Println("Visiting:", work)
				resp, err := client.Get(work)
				// check response
				if err != nil {
					id.errors.WriteString(work + "|" + err.Error() + "\n")
				} else {
					id.visited.WriteString(work + "|" + resp.Status + "\n")
					// get response body as string
					bodyBytes, err := ioutil.ReadAll(resp.Body)
					if err == nil {
						id.platformMatches(work, string(bodyBytes))
					}
				}
			}
		}
	}()
}