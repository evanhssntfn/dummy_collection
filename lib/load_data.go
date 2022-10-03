  package main

  import (
    "context"
  	"encoding/json"
  	// "net/http"
  	"strconv"
  	"strings"
  	"github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/esapi"
    "io/ioutil"
    "fmt"
    "os"
    //"log"
  )

  //data struct for post type

  type Data struct {
    Data []DataEntry `json:"data"`
  }

  type DataEntry struct {
    userId int `json:"userId"`
    id int `json:"id"`
    title string `json:"title"`
    body string `json:"body"`
  }

  func LoadData(es (*elasticsearch.Client)) {

    jsonFile, err := os.Open("data.json")

    if err != nil {
      fmt.Println(err)
    }

    fmt.Println("Successfully Opened data.json")

    defer jsonFile.Close()
		body, _ := ioutil.ReadAll(jsonFile)
    var result map[string]interface{}
    json.Unmarshal([]byte(body), &result)
    var final = result["data"]
    //fmt.Println(final)

    // var data Data
    // json.Unmarshal(body, &data)


    for _, v := range final.([]interface{}){
      val := v.(map[string]interface{})
      val["id"] = int(val["id"].(float64))
      uid := strconv.Itoa(val["id"].(int))
  		jsonString, _ := json.Marshal(v)
      fmt.Printf("val: %s", jsonString)
      fmt.Println("val:", uid)
  		request := esapi.IndexRequest{Index: "stsc", DocumentID: uid, Body: strings.NewReader(string(jsonString))}
  		request.Do(context.Background(), es)
	   }

    // for i := 0; i < len(data.Data); i++{
    //   uid := strconv.Itoa(data.Data[i].id)
  	// 	jsonString, _ := json.Marshal(data.Data[i])
    //   fmt.Printf("val: %s", jsonString)
    //   fmt.Println("val:", uid)
  	// 	// request := esapi.IndexRequest{Index: "stsc", DocumentID: uid, Body: strings.NewReader(string(jsonString))}
  	// 	// request.Do(context.Background(), es)
	  //  }

}
    // for k, v := range res {
    //   switch c := v.(type) {
    //   case string:
    //     fmt.Printf("Item %q is a string, containing %q\n", k, c)
    //   case float64:
    //     fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
    //   default:
    //     fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
    //   }
    // }







    //fmt.Println(posts)
    //return json.NewDecoder(r.Body).Decode(posts)
