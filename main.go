package main

import (
	"github.com/IrekRomaniuk/phcsv/utils"
	"fmt"
	"os"
	"log"
	"flag"	
)

var (
	//URL of Phantom endpoint
	URL = flag.String("url", "https://10.34.1.110/rest/decided_list/", "Phantom REST endpoint")
	//NAME of Phantom list
	NAME= flag.String("name", "Shields", "Name of custom list")
	//PATH to csv file
	PATH = flag.String("path", "./", "path to read csv file from")	
	//USER Phantom
	USER    = flag.String("u", "admin", "Phantom username")
	//PASS Phantom
	PASS    = flag.String("p", "", "Phantom password")
	//MIN Minimal number of lines in file
	MIN    = flag.Int("min", 0, "Min of lines in csv files (or will skip)")
	version = flag.Bool("v", false, "Prints current version")
	// Version : Program version
	Version = "No Version Provided" 
	// BuildTime : Program build time
	BuildTime = ""
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All jdk-rights reversed.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}	
}

//go run main.go -p='password' -name="Shields" -path="sample/sample.csv"
func main() {
	list := &utils.List{Content: [][]string{}}
	err := list.ReadFile(*PATH, *NAME)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("List size: %v\n", len(list.Content))
	if len(list.Content) > *MIN {
		fmt.Printf("url: %s\n",*URL + *NAME )
		ID, err := utils.PostPage(*URL + *NAME, *USER, *PASS, list)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Response ID: %v\n", ID)
	} else {
		fmt.Printf("Number of lines in csv file less than min: %v\n", *MIN)
	} 
}