package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func run(arg string) {
	cmd := exec.Command(arg)
	cmd.Start()
	cmd.Wait()
}

func main() {
	log.Println("Started go cmdrunner")
	run(strings.Join(os.Args, " "))
	log.Println("Finished go cmdrunner")
}
