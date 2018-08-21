package templates

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

func getClient() (*cloudformation.CloudFormation, error) {
	var client *cloudformation.CloudFormation

	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.SharedCredentialsProvider{
				Filename: "",
				Profile:  "frontend",
			},
			&credentials.EnvProvider{},
		})

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: creds,
	})
	if err != nil {
		return client, err
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return client, err
	}

	return cloudformation.New(sess), nil
}

func TestTemplate(t *testing.T) {
	client, _ := getClient()
	_, err := client.ValidateTemplate(&cloudformation.ValidateTemplateInput{
		TemplateBody: &Cfn,
	})

	if err != nil {
		t.Errorf("Cloudformation template invalid, %s", err.Error())
	}
}
