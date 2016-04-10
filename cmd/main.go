package main

import (
    "log"
    
    "github.com/learnaddict/dnsoverhttps"
)

func main() {
    log.Println("Question: A record for google.com")
    a, _ := dnsoverhttps.A("google.com")
    log.Printf("Answer: %s", a)
}