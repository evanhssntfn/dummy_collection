package main

import (
    "bufio"
    "fmt"
    "context"
    "os"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/esapi"
    "encoding/json"
    "bytes"
)

var es, _ = elasticsearch.NewDefaultClient()

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) load data")
		fmt.Println("2) get data")
		fmt.Println("3) search data by key val")
		option := ReadText(reader, "enter option")
		if option == "0" {
			Exit()
		} else if option == "1" {
			LoadData(es)
		} else if option == "2" {
			Get(reader)
		} else if option == "3" {
			Search(reader, "match")
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func Get(reader *bufio.Scanner) {
	id := ReadText(reader, "Enter datapoint ID")
	request := esapi.GetRequest{Index: "stsc", DocumentID: id}
	response, _ := request.Do(context.Background(), es)
	var results map[string]interface{}
	json.NewDecoder(response.Body).Decode(&results)
	fmt.Print(results["_source"].(map[string]interface{}))
}

func PrintData(dataentry map[string]interface{}) {
  userId := dataentry["userId"]
  id := dataentry["id"]
  title := dataentry["title"]
  body := dataentry["body"]

  fmt.Println(userId, id, title, body)
}


//thank god for the internet - I struggle so much with type assertions in go
func Search(reader *bufio.Scanner, querytype string) {
	key := ReadText(reader, "Enter key")
	value := ReadText(reader, "Enter value")
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			querytype: map[string]interface{}{
				key: value,
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	response, _ := es.Search(es.Search.WithIndex("stsc"), es.Search.WithBody(&buffer))
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		s := hit.(map[string]interface{})["_source"].(map[string]interface{})
		fmt.Println(s)
	}
}

func Exit() {
	os.Exit(0)
}

func ReadText(reader *bufio.Scanner, prompt string) string {
	fmt.Print(prompt + ": ")
	reader.Scan()
	return reader.Text()
}
