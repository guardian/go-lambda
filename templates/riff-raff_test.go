package templates

import (
	"testing"

	"github.com/andreyvit/diff"

	"github.com/guardian/go-lambda/config"
)

var expected string = `stacks:
- frontend
regions:
- eu-west-1

deployments:
  cfn:
    type: cloud-formation
    parameters:
      cloudFormationStackName: com-gu-frontend-test
      templatePath: cfn.yaml
      cloudFormationStackByTags: false
      templateParameters:
        VPC: vpc-5953433c
        Bucket: com-gu-frontend-golambda
        Key: test.zip
        Main: main
        Subnets: subnet-2324d37a,subnet-2419b341,subnet-e18d5e96
        ManagedPolicies: arn:a,arn:b
  lambda:
    type: aws-lambda
    parameters:
      functions:
        PROD:
          name: TODO
          filename: test.zip
      bucket: com-gu-frontend-test
    dependencies:
    - cfn
`

func TestConfig(t *testing.T) {
	conf := config.Config{
		Name:     "test",
		Profile:  "frontend",
		VpcID:    "vpc-5953433c",
		Subnets:  []string{"subnet-2324d37a", "subnet-2419b341", "subnet-e18d5e96"},
		Policies: []string{"arn:a", "arn:b"},
	}

	actual, err := RiffRaffConfig(conf)

	if err != nil {
		t.Errorf("config generation failed with error %s", err.Error())
		return
	}

	if actual != expected {
		t.Errorf("config did not match expected, %v", diff.LineDiff(actual, expected))
	}
}
