package helper

import (
	"strings"

	"github.com/plivox/gulmy/pkg/shell"
)

func GitCommitHash() string {
	version := shell.Cmd("git", "describe", "--long", "--tags", "--always").Output()
	return strings.TrimSpace(version)
}
