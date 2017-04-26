package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]
	argsString := strings.Join(argsWithoutProg, " ")

	if len(argsWithoutProg) < 2 {
		fmt.Printf("lazygit: %s is not a valid lazygit command.\nUsage: lazygit commit_message", argsString)
		return
	}

	cmd := exec.Command("git", "add", "-A")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("git add -A")

	cmd = exec.Command("git", "add", "-u")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("git add -u")

	cmd = exec.Command("git", "commit", "-m", argsString)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("git commit -m \"%s\"", argsString)

	cmd = exec.Command("git", "push")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("git push")

	fmt.Printf("Major success")
	return
}
