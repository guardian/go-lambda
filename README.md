
# Go Lambda

*WIP - much of the functionality is not there yet.**

    $ go install github.com/guardian/go-lambda

(Though see 'Installing Go' below if you haven't already done so
before running this command.)

Handles Guardian-related scaffolding for AWS lambda functions written
in Go.

* generate a lambda skeleton project
* generate a RiffRaff artifact, including code and relevant
  cloudformation

A lambda.json file is used for configuration.

Main commands:

    $ go-lambda new
    $ go-lambda build
    $ go-lambda help

`new` builds the skeleton project. It will optionally create a
Teamcity project for you.

`build` will generate the RiffRaff artifact including
cloudformation. This is useful for local testing, but typically will
be run on your build server.

## Installing Go

To install Go on Mac, run:

	$ brew install golang
	$ echo 'GOPATH=~/gocode' >> ~/.bashrc
	$ echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc

$GOPATH is used to store and locate dependencies.

**You should locate all projects under this path for now.**

Typically, under `$GOPATH/src/github.com/guardian`.

Go Modules solve this and are much better. However, they are *very*
new and VS Code doesn't support them yet unfortunately, though this is
expected imminently (in the next month).
