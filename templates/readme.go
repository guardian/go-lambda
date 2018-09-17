package templates

var Readme = `
# Go lambda function

Simple skeleton for an AWS (Golang) lambda function.

## Testing

A simple main_test.go file has been provided, which you can run using:

	$ go test

Read more about the Go testing package here:
https://golang.org/pkg/testing/. Spoiler: it's pretty basic.

## Useful packages

* [HTTP](https://golang.org/pkg/net/http/)
* [JSON](https://golang.org/pkg/encoding/json/)
* [Logging](https://golang.org/pkg/log/)

Strive to avoid adding dependencies outside of the standard library.

One possible exception is error handling. Because lambdas are simple
you can probably just log errors, but if you need something a bit
richer, I'd recommend:

https://github.com/pkg/errors

## Using the AWS SDK

You probably want to go get the AWS SDK as you'll likely use some of
it for your lambda:

	$ go get -u github.com/aws/aws-sdk-go/...

(In production, AWS provides this dependency for you.)

You can then import specific packages like so:

	$ import "github.com/aws/aws-sdk-go/service/s3"

All of the clients are set up in a similar way:

	import (
		"github.com/aws/aws-sdk-go/aws/credentials"
		"github.com/aws/aws-sdk-go/aws/session"
	)

	// build a credentials provider
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.SharedCredentialsProvider{
				Filename: "",
				Profile:  "frontend",
			},
			&credentials.EnvProvider{},
		})

	// grab credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: creds,
	})

	// check if they were found
	_, err = sess.Config.Credentials.Get()


Note, in actual code you should do something with the errors (probably
just log.Fatal...).  `
