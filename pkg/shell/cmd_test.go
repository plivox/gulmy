package shell

import (
	"testing"
)

func TestShellCmd(t *testing.T) {
	MakeStyle()
	Cmd("ls", "-l").Dir("/tmp").Run()
}
