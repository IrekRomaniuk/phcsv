package utils

import (
	"testing"
	"log"	
	"fmt"
	"flag"
)
var (
	user = flag.String("u", "", "user")
	pass = flag.String("p", "", "pass")
)
//TestReadFile  : go test -run TestReadFile
func TestReadFile(t *testing.T) {
	list := &List{Content: [][]string{}}
	//fmt.Printf("List name: %v\n", list.Name)		
	err := list.ReadFile("../sample/sample.csv", "Shields") 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Content: %v\n", *list)
}
//TestPostPage : go test -run TestPostPage -args -u=user -p=pass
func TestPostPage(t *testing.T) {
	list := List{Content: [][]string{{"1.1.1.2","Testphcsv"},{}}, Name: "phcsv2"}
	fmt.Printf("List: %v\n", list)
	ID, err := PostPage("https://10.34.1.110/rest/decided_list", *user, *pass, list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v\n", ID)
}