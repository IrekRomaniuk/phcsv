package utils
//go test -run TestReadFile
import (
	"encoding/csv"
	"os"
	"bufio"
	"fmt"
	"net/http"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
)

//List https://my.phantom.us/3.0/docs/rest/lists
type List struct {
	Content [][]string `json:"content"`
	Name    string     `json:"name"`
}

//ReadFile csv into List name and return encoded into json list
func (list *List) ReadFile(file, name string) error {
	f, err := os.Open(file)
	if err != nil {
		return err 
	}
	r := csv.NewReader(bufio.NewReader(f))	
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}
	fmt.Printf("Lines: %v\n", len(lines)) 
    /*for i := range lines {   		
		fmt.Println(lines[i])         	
    }*/
	*list = List{Content: lines, Name: name} 
	//copy(*list.Content, lines)
	return nil
}
//Response from Phantom 
type Response struct {
	ID int64 `json:"id"`
	Success bool `json:"success"` 
}
//PostPage tp Phantom and return response json with "id" and "success" 
func PostPage(url, user, pass string, data interface{}) (int64, error) {
	var response Response
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(data)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return 0, err
	}
	req.SetBasicAuth(user, pass)
	//req.Header.Set("ph-auth-token", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	err = json.Unmarshal(htmlData, &response)
	if err != nil {
		return 0, err
	}
	return response.ID, nil	
}