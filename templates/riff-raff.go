package templates

import (
	"encoding/json"
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
      cloudFormationStackName: {{.Name}}
      templatePath: cfn.yaml
      cloudFormationStackByTags: false
  lambda:
    type: aws-lambda
    parameters:
      functions:
        PROD:
          name: frontend-dokr-PROD-DokrLambda-1F1WWBIGRY1LJ
          filename: {{.Name}}.zip
      bucket: com-gu-{{.Profile}}-{{.Name}}
    dependencies:
    - cfn
`

func RiffRaffConfig(conf config.Config) (string, error) {
	tpl, err := template.New("rr").Parse(rr)
	if err != nil {
		return "", err
	}

	var buffer strings.Builder
	err = tpl.Execute(&buffer, conf)
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
