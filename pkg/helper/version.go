package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func VersionFromFile() string {
	buf, err := ioutil.ReadFile("VERSION")
	if err != nil {
		log.Fatalf("error reading file VERSION: %v\n", err)
	}
	return strings.TrimSpace(string(buf))
}

func VersionFromPackageJson() string {
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	reader, err := os.Open("package.json")
	if err != nil {
		log.Fatal("Failed to open package.json")
	}
	defer reader.Close()

	config := map[string]interface{}{}
	if err := json.NewDecoder(reader).Decode(&config); err != nil {
		log.Fatal("Failed to decode package.json")
	}
	return config["version"].(string)
}
