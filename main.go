package main

import (
	"log"
	"os/exec"
	"os"
)

func main() {
	cmd := exec.Command("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe", "https://www.youtube.com/watch?v=dQw4w9WgXcQ") // execute edge in headless mode

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}