package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

const sockAddr = "/tmp/cmdrunner.sock"

func run(arg string) {
	if len(arg) == 0 {
		log.Println("ERROR: Invalid arg")
		return
	}
	log.Printf("running *%v*\n", arg)
	cmd := exec.Command("bash", "-c", arg)
	out, err := cmd.Output()
	if err != nil {
		log.Println("ERROR: ", err)
	}
	log.Println(string(out))
}

func handle(c net.Conn) {
	remoteAddr := c.RemoteAddr()
	defer log.Println("connection closed for ", remoteAddr)
	defer c.Close()
	r := bufio.NewReader(c)
	for {
		rawline, err := r.ReadString('\n')

		//connection has probably closed
		if err != nil {
			log.Println(err)
		} else {
			cmd := strings.Trim(rawline, "\n\r")
			run(cmd)
		}
	}
}

func main() {
	l, err := net.Listen("unix", sockAddr)
	defer l.Close()

	if err != nil {
		log.Fatal("Failed to open socket", err)
	}
	log.Println("Started go cmdrunner at", sockAddr)
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("error in accept", err)
			break
		}
		log.Println("accepted connection")
		go handle(c)
	}
	log.Println("Finished go cmdrunner")
}
