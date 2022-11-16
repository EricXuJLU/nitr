package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "nsenter -t 1 sh -c \"ps -aux\"")
	op, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	out, err := io.ReadAll(op)
	op.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
