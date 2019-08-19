package shell

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type cmd struct {
	// name      string
	// args      []string
	makeStyle bool
	cmd       *exec.Cmd
}

func Cmd(name string, args ...string) *cmd {
	c := exec.Command(name, args...)
	// c.Stdout = os.Stdout
	// c.Stderr = os.Stderr
	// return &cmd{name, args, GlobalMakeStyle, c}
	return &cmd{GlobalMakeStyle, c}
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

	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	if err := c.cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func (c *cmd) Output() string {
	out, err := c.cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
