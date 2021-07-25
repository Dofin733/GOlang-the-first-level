package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
 Port string
 Name string
 Count int
 Db_url string
 Jaeger_url string
 Sentry_url string
 Kafka_broker string
}

func main () {
	var configFileName = flag.String("config", "", "Config file")
	flag.Parse()

	configData, err := os.ReadFile(*configFileName)
	if err != nil {
		log.Fatalln(err)
	}
	config := Config{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
}

