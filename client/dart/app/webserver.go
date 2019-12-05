package main

import (
	"encoding/json"
	"fmt"
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

	c := &config{
		ChannelAddress: endpoint,
	}
	f, err := os.Open("build/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(f).Encode(&c)
}

func main() {
	log.Fatal(http.ListenAndServe(address, http.FileServer(http.Dir("./build"))))
}