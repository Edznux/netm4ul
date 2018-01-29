package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("dirb", "http://example.com") //Exec
	var out bytes.Buffer //Data buffer
	cmd.Stdout = &out
	err := cmd.Run() //Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(out.String()) //Print result
}