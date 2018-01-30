package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"regexp"
)

func main() {
	fmt.Println("Test affichage")

	//Conf
	var hostname string = "http://example.com"
	var wordlist string = "/usr/share/wordlists/dirb/others/best15.txt"


	cmd := exec.Command("dirb", hostname, wordlist, "-S") //Execution
	var out bytes.Buffer //Data buffer
	cmd.Stdout = &out
	err := cmd.Run() //Error
	if err != nil {
		log.Fatal(err)
	}

	var result string = out.String()

	words := strings.Fields(result)
	
	re := regexp.MustCompile("http([a-z]+)*") //URL Regex

	for i, word := range words {
		if re.FindString(word) != "" {
			i++ //lol
			fmt.Println(word)
		}
	}
}

//type Dirb struct {
//	Config  ConfigToml
//	Result  []byte
//	Nmaprun *mynmap.NmapRun
//}