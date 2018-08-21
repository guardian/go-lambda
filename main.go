package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/guardian/go-lambda/templates"
)

type Statement struct {
	Effect   string
	Action   []string
	Resource string
}

type Config struct {
	Name     string
	VpcID    string
	Subnets  []string
	Policies []Statement
}

func main() {
	switch os.Args[1] {
	case "new":
		fmt.Println("Making a new project...")

		config := Config{}
		name := readLine("Project name")
		config.Name = name

		vpcID := readLine("VPC ID")
		config.VpcID = vpcID

		subnets := readLine("Subnets (comma-separated list")
		config.Subnets = strings.Split(subnets, ",")

		fmt.Println("Thanks! Writing project files...")

		err := mkdir(config.Name)
		check(err)

		err = writeConfig(fmt.Sprintf("%s/lambda.conf", config.Name), config)
		check(err)

		err = writeFile(fmt.Sprintf("%s/main.go", config.Name), []byte(templates.Lambda))
		check(err)

		fmt.Println("Done! You're ready to code :)")

	default:
		fmt.Println("Unrecognised command! Exiting...")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt, ": ")
	var input string
	fmt.Scanln(&input)
	return input
}

func mkdir(name string) error {
	return os.Mkdir(name, os.ModePerm)
}

func writeConfig(path string, config Config) error {
	s, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return writeFile(path, s)
}

func writeFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}
