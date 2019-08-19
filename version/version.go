package version

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	s "github.com/plivox/gulmy/shell"
)

func FromFile() string {
	buf, err := ioutil.ReadFile("VERSION")
	if err != nil {
		log.Fatalf("error reading file VERSION: %v\n", err)
	}
	return strings.TrimSpace(string(buf))
}

func FromGit() string {
	version := s.Cmd("git", "describe", "--long", "--tags", "--dirty", "--always").Output()
	return strings.TrimSpace(version)
}

func FromPackageJson() string {
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
