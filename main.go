package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/guardian/go-lambda/config"
	"github.com/guardian/go-lambda/templates"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Must provide command to run. See 'help' for more info.")
	}

	switch os.Args[1] {
	case "help":
		helpMessage := `
go-lambda [cmd]
  new           - builds the skeleton project
  build         - generates RiffRaff artifact including cloudformation
  create-lamdba - creates the lambda in AWS
  help          - provides help information
		`
		log.Print(helpMessage)
	case "new":
		// TODO better to just use actual go package here stored in a separate
		// github. It will be easier to test and iterate on.

		// This tool should really be about bringing things together and the
		// ops side.
		conf := config.Config{}

		name := readLine("Project name")
		conf.Name = name

		VCS := readLine("VCS URL ('git@...')")
		conf.VCSURL = VCS

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

		err = writeFile(fmt.Sprintf("%s/main_test.go", conf.Name), []byte(templates.LambdaTest))
		check(err)

	case "build":
		fmt.Println("Warning: this command is currently buggy.")

		conf, err := config.Read("lambda.conf")
		check(err)

		err = rmDir("target")
		check(err)

		err = mkdir("target/cfn")
		check(err)
		err = mkdir("target/lambda")
		check(err)

		rr, err := templates.RiffRaffConfig(conf)
		check(err)
		err = writeFile("target/riff-raff.yaml", []byte(rr))
		check(err)

		err = writeFile("target/cfn/cfn.yaml", []byte(templates.Cfn))
		check(err)

		err = exec.Command("go", "build", "-o", "target/lambda/lambda.go", "main.go").Run()
		check(err) // TODO can include stderr here

		buildJSON, err := templates.BuildJSON(conf)
		check(err)
		err = writeFile("target/build.json", []byte(buildJSON))

		// TODO use Go AWS SDK
		// S3ProjectTarget := fmt.Sprintf("s3://riffraff-artifact/%s/1", conf.ProjectName())
		// S3BuildTarget := fmt.Sprintf("s3://riffraff-builds/%s/1/build.json", conf.ProjectName())
		// err = exec.Command("aws", "s3", "cp", "--profile", "deployTools", "--acl", "bucket-owner-full-control", "--region=eu-west-1", "--recursive", "target", S3ProjectTarget).Run()
		// check(err)
		// err = exec.Command("aws", "s3", "cp", "--profile", "deployTools", "--acl", "bucket-owner-full-control", "--region=eu-west-1", "--recursive", "target/build.json", S3BuildTarget).Run()
		// check(err)

	case "create-lambda":
		// TODO this will create the lamdba in AWS and update your lamdba.conf accordingly
		fmt.Println("This command is not yet implemented sorry.")

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
	return os.MkdirAll(name, os.ModePerm)
}

func writeFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}

func rmDir(path string) error {
	return os.RemoveAll(path)
}
