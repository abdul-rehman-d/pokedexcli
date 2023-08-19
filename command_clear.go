package main

import (
	"os"
	"os/exec"
)

func callbackClear(_ *config, _ ...string) error {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	return c.Run()
}
