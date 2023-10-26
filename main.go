package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("begin")

	fmt.Println("output environment variables")
	for _, kv := range os.Environ() {
		fmt.Println(kv)
	}

	fmt.Println("output to GITHUB_OUTPUT")
	gitHubOutput := os.Getenv("GITHUB_OUTPUT")
	fooInput := os.Getenv("FOO_INPUT")

	output(gitHubOutput, fooInput)

	fmt.Println("end")
}

func output(gitHubOutput, fooInput string) {
	f, err := os.Create(gitHubOutput)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", gitHubOutput, err)
	}
	defer f.Close()

	s := fmt.Sprintf("hoge-output=Your-input-was-%s-is-it-right?", fooInput)
	if _, err := f.WriteString(s); err != nil {
		log.Fatalf("failed to write string: %v", err)
	}
}
