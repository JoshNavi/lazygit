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
	log.Printf("git add -A")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "add", "-u")
	log.Printf("git add -u")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "commit", "-m", argsString)
	log.Printf("git commit -m \"%s\"", argsString)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Check if you have changes to commit. Error: %s", err)
	}

	cmd = exec.Command("git", "push")
	log.Printf("git push")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Major success")
	return
}
