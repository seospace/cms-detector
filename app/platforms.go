package main

import (
	"os"
	"io/ioutil"
	"strings"
)

var fpFolder = "./footprints/" // has to end with "/"
var idFolder = "./identified/" // has to end with "/"

type Platform struct {
	name string
	file *os.File
	footprints []string
}

// parses footprints in fpFolder and returns map of footprints + map of files
func ParseFootprints() ([]Platform)  {
	var platforms []Platform
	// read all files in directory
	files, err := ioutil.ReadDir(fpFolder)
	if err != nil { panic(err) }
	// iterate over each file in dir
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			fName := f.Name()         // platform file name
			fPath := fpFolder + fName // platform file path
			fBytes, _ := ioutil.ReadFile(fPath)
			footprints := strings.Split(string(fBytes), "\n")
			if len(footprints) >= 1 {
				// trim each footprint
				for i, fp := range footprints {
					footprints[i] = strings.TrimSpace(fp)
				}
				// get platform name
				pName := strings.Replace(fName, ".txt", "", -1) // platform name
				// finally set what we need
				p := Platform{
					name:       pName,
					file:       getIdentifiedFile(pName),
					footprints: footprints,
				}
				platforms = append(platforms, p)
			}
		}
	}
	return platforms
}

// returns file where we save identified urls for pName (platform name)
func getIdentifiedFile(pName string) *os.File {
	fPath := idFolder + pName + ".txt"
	f, err := os.OpenFile(fPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil { panic(err) }
	return f
}