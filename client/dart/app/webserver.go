package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	portKey = "PORT"
	endpointKey = "ENDPOINT"
)

var (
	address, endpoint string
)

type config struct {
	ChannelAddress string `json:"channelAddress"`
}

func init() {
	requiredEnv := map[string]struct{}{
		portKey: {},
		endpointKey: {},
	}

	for k := range requiredEnv {
		if os.Getenv(k) == "" {
			log.Fatalf("%s empty but expected from environment variables", k)
		}
	}

	address = fmt.Sprintf(":%s", os.Getenv(portKey))
	endpoint = os.Getenv(endpointKey)
	r := regexp.MustCompile("^http")
	if !r.MatchString(endpoint) {
		log.Fatalf("%s does not match required pattern ^http://", endpoint)
	}
}

func main() {
	c := &config{
		ChannelAddress: endpoint,
	}
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	configData := fmt.Sprintf("window.config = %s", string(data))
	err = ioutil.WriteFile("./build/config.js", []byte(configData), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(address, http.FileServer(http.Dir("./build"))))
}