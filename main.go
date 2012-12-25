package main

import (
	"log"
	"os"
	"os/exec"
)

func run(arg ...string) {
	//cmd := exec.Command("bash", "-c", arg)
	log.Printf("running *%v*\n", arg)
	cmd := exec.Command(arg[0], arg[1:]...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("ERR", err)
	}
	log.Println(string(out))
}

func main() {
	log.Println("Started go cmdrunner")
	log.Println(os.Args)
	run(os.Args[1:]...)
	log.Println("Finished go cmdrunner")
}
