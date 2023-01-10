package helper

import (
	"fmt"
	"testing"
)

func TestVersionFromFile(t *testing.T) {
	fmt.Println(VersionFromFile())
}

func TestVersionFromPackageJson(t *testing.T) {
	fmt.Println(VersionFromPackageJson())
}
