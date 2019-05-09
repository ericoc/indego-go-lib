package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func main() {

    // Set URL of the Philadelphia Indego Bike Share API, along with a user-agent to use for the HTTP request
    url := "https://www.rideindego.com/stations/json/"
    user_agent := "Indego Go API Library - https://github.com/ericoc/indego-go-lib"

    // Perform the HTTP request and get the resulting JSON
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("User-Agent", user_agent)
    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    // Do something that I do not understand at all with the JSON
    var result map[string]interface{}
    json.Unmarshal([]byte(body), &result)

    // Print whatever output I got from the stuff above that I don't quite understand
    fmt.Printf("Result: %s", result["features"])
}
