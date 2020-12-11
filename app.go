package main

import (
	"fmt"
	"os"
    "strings"
    "github.com/kelseyhightower/envconfig"
    "net/http"
    "io/ioutil"
)

type DTOptions struct {
  BaseURL string
  APIToken string
  MetricsList string
  Delimiter string
}

func main() {

  var delimiter = "#" // Default delimiter is a hash. This can be overwritten with DT_DELIMITER=***
  // Parse input options
  var dtoptions DTOptions
  error := envconfig.Process("dt", &dtoptions)
  
  if (error != nil) {
    fmt.Println("got an error")
  }
  
  // These params are mandatory.
  if (dtoptions.BaseURL == "") {
    fmt.Println("NO DT_BASEURL environment variable defined.")
    os.Exit(1)
  }
  if (dtoptions.APIToken == "") {
    fmt.Println("NO DT_APITOKEN environment variable defined.")
    os.Exit(1)
  }
  if (dtoptions.MetricsList == "") {
  fmt.Println("NO DT_METRICSLIST environment variable defined.")
    os.Exit(1)
  }
  // if we have a custom delimiter, set it now.
  if (dtoptions.Delimiter != "") {
    delimiter = dtoptions.Delimiter
  }
  
  fmt.Println("Delimiter:", delimiter)
  
  var metricsList = strings.Split(dtoptions.MetricsList, delimiter)
  

  dynatrace_api_url := dtoptions.BaseURL + "/api/v2/metrics/ingest"
  
  
  // Build body. Metrics are passed one per line.
  var bodyContent = ""
  for _, metric := range metricsList {
    bodyContent += metric + "\n"
  }
  
  fmt.Println("============================")
  fmt.Println("Pushing Metrics to Dynatrace")
  fmt.Println("============================")
  
  // Create HTTP Client
  client := &http.Client{}
  
  req,_ := http.NewRequest("POST", dynatrace_api_url, strings.NewReader(bodyContent))
  req.Header.Add("Authorization", "Api-Token " + dtoptions.APIToken)

  // Send Request
  resp,_ := client.Do(req)
  defer resp.Body.Close()
    
  fmt.Println("Response Code:", resp.StatusCode)
  
  body,_ := ioutil.ReadAll(resp.Body)
 
  fmt.Printf("Body: %s", body)

}
