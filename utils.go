package main

import (
	"strings"
	"os"
	"time"
)

func (id *Identifier) platformMatches(u, body string) {
	for _, platform := range id.platforms {
		for _, footprint := range platform.footprints {
			if strings.Index(body, footprint) != -1 {
				platform.file.WriteString(u + "\n")
				break
			}
		}
	}
}


func GetVisitAndErrorFiles() (*os.File, *os.File){
	vP, err := os.OpenFile("./results/visited.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil { panic(err) }
	eP, err := os.OpenFile("./results/error.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil { panic(err) }
	return vP, eP
}

func ConvertTimeout(t int) time.Duration {
	timeoutDuration := time.Duration(t)
	return time.Duration(timeoutDuration * time.Second)
}