package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Session aws session to access services
var Session *session.Session

func init() {
	Session = session.Must(session.NewSession(&aws.Config{
		S3DisableContentMD5Validation: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
		EndpointResolver:              endpoints.ResolverFunc(serviceProxy),
	}))
}

func serviceProxy(
	service string,
	region string,
	optFns ...func(*endpoints.Options),
) (
	resolver endpoints.ResolvedEndpoint,
	err error,
) {
	if service := serviceAddress(service); service != "" {
		resolver = endpoints.ResolvedEndpoint{URL: service}
	} else {
		resolver, err = endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	return
}

func serviceAddress(service string) (address string) {
	var port string
	switch service {
	case dynamodb.EndpointsID:
		port = os.Getenv("LOCALSTACK_DYNAMO")
	case s3.EndpointsID:
		port = os.Getenv("LOCALSTACK_S3")
	case lambda.EndpointsID:
		port = os.Getenv("LOCALSTACK_LAMBDA")
	}

	if host := os.Getenv("LOCALSTACK_HOST"); host != "" && port != "" {
		address = fmt.Sprintf("http://%s:%s", host, port)
	}

	return
}
