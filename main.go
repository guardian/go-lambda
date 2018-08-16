package main

import (
	"fmt"
	"os"
	"strings"
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

		fmt.Println("Almost done!")
		fmt.Println("You can add AWS permissions in the lambda.conf file at the root of the project folder.")
		fmt.Println("Your project is ready for coding! :)")

		fmt.Println(config)
	default:
		fmt.Println("Unrecognised command! Exiting...")
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt, ": ")
	var input string
	fmt.Scanln(&input)
	return input
}
