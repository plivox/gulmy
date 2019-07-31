package shell

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type cmd struct {
	makeStyle bool
	cmd       *exec.Cmd
}

func (c *cmd) Dir(dir string) *cmd {
	c.cmd.Dir = dir
	return c
}

func (c *cmd) MakeStyle() *cmd {
	c.makeStyle = true
	return c
}

func (c *cmd) Run() {
	if c.makeStyle || GlobalMakeStyle {
		log.Println(c.cmd.Args[0], strings.Join(c.cmd.Args[1:], " "))
	}

	if err := c.cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func Cmd(name string, args ...string) *cmd {
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return &cmd{GlobalMakeStyle, c}
}
