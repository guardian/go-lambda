package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/guardian/go-lambda/aws"
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
  build         - generates and uploads RiffRaff artifact including cloudformation
  create-lamdba - creates the lambda in AWS
  help          - provides help information
		`
		log.Print(helpMessage)
	case "new":
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

		// TODO what if dependencies not found?
		err = exec.Command("go", "build", "-o", "target/lambda/lambda.go", "main.go").Run()
		check(err) // TODO can include stderr here

		buildJSON, err := templates.BuildJSON(conf)
		check(err)
		err = writeFile("target/build.json", []byte(buildJSON))

		client, err := aws.GetClient()
		check(err)

		// TODO use build number here instead of 1
		S3KeyPrefix := fmt.Sprintf("/%s/1/", conf.ProjectName())
		paths := lsDir("target")
		pathKeys := PathsToKeys(paths, S3KeyPrefix, "target/")
		aws.UploadFiles(client, "riffraff-artifact", pathKeys)

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

func isDirectory(path string) bool {
	fd, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	switch mode := fd.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		return false
	}
	return false
}

func lsDir(path string) []string {
	fileList := []string{}
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if isDirectory(path) {
			return nil
		} else {
			fileList = append(fileList, path)
			return nil
		}
	})

	return fileList
}

// Build map of file paths to S3 keys for Riffraff upload
func PathsToKeys(
	paths []string,
	keyPrefix string,
	stripPrefix string,
) map[string]string {

	m := map[string]string{}
	for _, path := range paths {
		key := keyPrefix + strings.TrimPrefix(path, stripPrefix)
		m[path] = key
	}

	return m
}
