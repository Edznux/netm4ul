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

//Function what Remove duplicate
func unique(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{} 
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func main() {
	//Conf
	var hostname string = "https://example.com"
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

    uniqueSlice := unique(wordList)
	dirbStruct := Dirb{Host : hostname,Urls : uniqueSlice}
	fmt.Println(dirbStruct)

}