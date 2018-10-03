package templates

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/guardian/go-lambda/config"
)

var rr string = `stacks:
- {{.Profile}}
regions:
- eu-west-1

deployments:
  cfn:
    type: cloud-formation
    parameters:
      cloudFormationStackName: com-gu-{{.Profile}}-{{.Name}}
      templatePath: cfn.yaml
      cloudFormationStackByTags: false
      templateParameters:
        VPC: {{.VpcID}}
        Bucket: {{.Bucket}}
        Key: {{.Key}}
        Main: {{.Main}}
        Subnets: {{StringsJoin .Subnets ","}}
        ManagedPolicies: {{StringsJoin .Policies ","}}
  lambda:
    type: aws-lambda
    parameters:
      functions:
        PROD:
          name: {{.FunctionName}}
          filename: {{.Name}}.zip
      bucket: com-gu-{{.Profile}}-{{.Name}}
    dependencies:
    - cfn
`

type rRConfig struct {
	config.Config
	Key          string
	Bucket       string
	Main         string
	FunctionName string
}

func RiffRaffConfig(conf config.Config) (string, error) {
	bucket := fmt.Sprintf("com-gu-%s-golambda", conf.Profile)
	key := conf.Name + ".zip"

	tpl, err := template.New("rr").Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(rr)
	if err != nil {
		return "", err
	}

	rrConfig := rRConfig{
		Config:       conf, // embedded
		Key:          key,
		Bucket:       bucket,
		Main:         "main",
		FunctionName: "TODO",
	}

	var buffer strings.Builder
	err = tpl.Execute(&buffer, rrConfig)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

type meta struct {
	ProjectName string `json:"projectName"`
	BuildNumber string `json:"buildNumber"`
	StartTime   string `json:"startTime"`
	Revision    string `json:"revision"`
	VCSURL      string `json:"vcsURL"`
	Branch      string `json:"branch"`
}

func BuildJSON(conf config.Config) (string, error) {
	now, _ := time.Now().MarshalText()

	// TODO use env vars for the rest
	m := meta{
		ProjectName: conf.ProjectName(),
		BuildNumber: "DEV",
		StartTime:   string(now),
		Revision:    "1",
		VCSURL:      conf.VCSURL,
		Branch:      "master",
	}

	bytes, err := json.Marshal(m)

	return string(bytes), err
}
