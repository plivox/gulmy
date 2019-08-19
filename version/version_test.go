package version

import (
	"fmt"
	"testing"
)

func TestVersionFromFile(t *testing.T) {
	fmt.Println(FromFile())
}

func TestVersionFromGit(t *testing.T) {
	fmt.Println(FromGit())
}
