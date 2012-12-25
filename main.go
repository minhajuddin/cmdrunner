package main

import (
	"log"
	"os/exec"
	"net"
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

	l, err := net.Listen("unix", "/tmp/cmdrunner.sock")

	if err != nil {
	  log.Fatal("Failed to open socket", err)
	}
	

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("error in accept", err)
			break
		}
		log.Println("accepted connection")
		go handle(c)
	}
	//log.Println(os.Args)
	//run(os.Args[1:]...)
	log.Println("Finished go cmdrunner")
}
