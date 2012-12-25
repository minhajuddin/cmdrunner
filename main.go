package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func run(arg string) {
	cmd := exec.Command("bash", "-c", arg)
	out, err := cmd.Output()
	if err != nil {
	  log.Fatal("ERR", err)
	}
	log.Println(">"+ arg)
	log.Println(string(out))
}

func main() {
	log.Println("Started go cmdrunner")
	run(strings.Join(os.Args[1:], " "))
	log.Println("Finished go cmdrunner")
}
