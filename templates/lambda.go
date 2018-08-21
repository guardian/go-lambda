package templates

// Lambda template
var Lambda = `
import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// See https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html

func main() {
	lambda.Start(handle)
}

func handle(ctx context.Context, event events.CloudWatchEvent) error {
	lc, _ := lambdacontext.FromContext(ctx)

	// Access event example
	log.Printf("Lambda invoked at %s", event.Time.String())

	// Access context example
	log.Printf("Env vars are: %v", lc.ClientContext.Env)

	return nil
}
`

func add(x int, y int) int {
	return x + y
}
