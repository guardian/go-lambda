package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/guardian/go-lambda/config"
	"github.com/guardian/go-lambda/templates"
)

func main() {
	switch os.Args[1] {
	case "new":
		conf := config.Config{}

		name := readLine("Project name")
		conf.Name = name

		profile := readLine("Janus profile")
		conf.Profile = profile

		vpcID := readLine("VPC ID")
		conf.VpcID = vpcID

		subnets := readLine("Subnets (comma-separated list")
		conf.Subnets = strings.Split(subnets, ",")

		err := mkdir(conf.Name)
		check(err)

		err = config.Write(fmt.Sprintf("%s/lambda.conf", conf.Name), conf)
		check(err)

		err = writeFile(fmt.Sprintf("%s/main.go", conf.Name), []byte(templates.Lambda))
		check(err)

	case "build":
		// will build a deployable artifact, including (currently)
		// - build binary
		// - generate cloudformation
		// - package in riffraff

		// A --publish flag (defaulting to false) controls writing
		// artifact to S3.

		// test cloudformation by trying to validate it

		// build binary
		// "go build main.go"
		conf, err := config.Read("lambda.conf")
		check(err)

		err = writeFile(fmt.Sprintf("%s/cfn.yaml", conf.Name), []byte(templates.Cfn))
		check(err)

		// package in riffraff
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

func writeFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}
