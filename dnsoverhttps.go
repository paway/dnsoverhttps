package dnsoverhttps

import ( 
    "fmt"
	//"io/ioutil"
	"log"
	"net/http"
    "encoding/json"
)

var (
    baseurl = "https://dns.google.com/"
    resolve = "resolve"
)



type DNSResponse struct {
  Status int32 `json:"Status"`
  TC bool `json:"TC"`
  RD bool `json:"RD"`
  RA bool `json:"RA"`
  AD bool `json:"AD"`
  CD bool `json:"CD"`
  Question []Question `json:"Question"`
  Answer []Answer `json:"Answer"`
  Additional []Additional `json:"Additional"`
  edns_client_subnet string `json:"edns_client_subnet"`
  Comment string `json:"Comment"`
}

type Question struct {
    Name string `json:"name"`
    TypeId int32 `json:"type"`
}


type Answer struct {
    Name string `json:"name"`
    TypeId int32 `json:"type"`
    TTL int32 `json:"TTL"`
    Data string `json:"data"`
}


type Additional struct {
    Unknown string
}

// Call the API request
func Call(url string) {
    log.Printf("http.Get: %s", url)
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    decoder := json.NewDecoder(res.Body)
    var data DNSResponse
    err = decoder.Decode(&data)
    if err != nil {
        panic(err)
    }
    for i, answer := range data.Answer{
        fmt.Printf("%d: %#v %#v\n", i, answer.Data, answer.Name)
    }
    fmt.Printf("%+v\n", data)
}

// A record query
func A(address string) (string, error) {
    log.Printf("Resolving A: %s", address)
    query := "https://dns.google.com/resolve?name=" + address
    Call(query)
    return "", nil
}

