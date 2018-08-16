# Go Lambda

*WIP - much of the functionality is not there yet.**

Handles Guardian-related scaffolding for AWS lambda functions written in Go.

* generate a lambda skeleton project
* generate a RiffRaff artifact, including code and relevant cloudformation

A lambda.json file is used for configuration.

Main commands:

    $ go-lambda new
    $ go-lambda build
    $ go-lambda help

`new` builds the skeleton project. It will optionally create a Teamcity project
for you.

`build` will generate the RiffRaff artifact including cloudformation. This is
useful for local testing, but typically will be run on your build server.
