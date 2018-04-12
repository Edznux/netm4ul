package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"regexp"
)

type Dirb struct {
	Host string
	Urls []string
}

//func (d Dirb) String() string {
//	return fmt.Sprintf(d.Host, d.Urls)
//}

func main() {
	fmt.Println("Test affichage")
	//Conf
	var hostname string = "example.com"
	var wordlist string = "/usr/share/wordlists/dirb/others/best15.txt" //Test dico

	//Command
	cmd := exec.Command("dirb", hostname, wordlist, "-S") //Execution
	var out bytes.Buffer //Data buffer
	var stderr bytes.Buffer //Data buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run() //Error

	if err != nil {
		log.Println(stderr.String())
		log.Fatal(err)
	}

	//result
	result := out.String()

	words := strings.Fields(result)
	
	re := regexp.MustCompile("https?://([a-z]+)*") //URL Regex
	var wordList []string
	for _, word := range words {
		if re.FindString(word) != "" {
			wordList = append(wordList, word)
			fmt.Println(word)
		}
	}

	dirbStruct := Dirb{Host : hostname,Urls : wordList}
	fmt.Println(dirbStruct)

}
