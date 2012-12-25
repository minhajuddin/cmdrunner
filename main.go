package main

import (
	"log"
	"os"
	"os/exec"
)

func run(args ...string) {
	if(len(args) == 0){
		log.Println("ERROR: Invalid number of args")
		return
	}
	log.Printf("running *%v*\n", args)
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	log.Println(string(out))
}

func main() {
	log.Println("Started go cmdrunner")
	log.Println(os.Args)
	run(os.Args[1:]...)
	log.Println("Finished go cmdrunner")
}
