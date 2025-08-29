package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	data, err := ioutil.ReadFile("rules/cre-2025-0071/autogpt-infinite-loop-detection.yaml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var result interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	fmt.Println("YAML is valid!")
}
