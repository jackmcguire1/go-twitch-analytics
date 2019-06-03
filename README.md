# go-twitch-analytics

[![Build Status](https://travis-ci.org/jackmcguire1/go-twitch-analytics.svg?branch=master)](hhttps://travis-ci.org/jackmcguire1/go-twitch-analytics)
[![Go Report Card](https://goreportcard.com/badge/github.com/jackmcguire1/go-twitch-analytics)](https://goreportcard.com/report/github.com/jackmcguire1/go-twitch-analytics)

[git]:    https://git-scm.com/
[golang]: https://golang.org/
[dlv]:    https://github.com/go-delve/delve
[modules]: https://github.com/golang/go/wiki/Modules
[goLand]: https://www.jetbrains.com/go/
[golint]: https://github.com/golangci/golangci-lint
[aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html
[aws-cli-config]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
[aws-sam-cli]: https://github.com/awslabs/aws-sam-cli
[localstack]: https://github.com/localstack/localstack

> This repo contains a severless stack example to ingest Twitch analytics daily onto S3

## ABOUT
Twitch recently removed the Game and Extension analytics webhooks,<br>
this resulted in me to develop an alternative <strong>serverless</strong> solution to fetch the statistics automatically.

### SERVERLESS STACK:
- Two lambda functions triggered by cloud watch scheduled events
- S3 Bucket
- User DynamoDB Table

This repo will deploy two serverless lambda functions to AWS, which will:<br>

- Renew a stored <strong>user access token</strong> within dynamoDB
- Call one of the <strong>Twitch</strong> extension/game analytics endpoints<br> 
- Store CSV file(s) onto S3

NOTE:
>  The two deployed lambda functions are configured to run at
- Fetch Extension Analytics -  <strong>02:00am GMT</strong> 
- Fetch Game  Analytics -  <strong>02:30am GMT</strong>
<br>

### Example S3 Locations
> s3://{{bucket-name}}/extension-analytics/overview_v2/ext-id/yyyy-mm-dd/*.csv

> s3://{{bucket-name}}/game-analytics/overview_v2/game-id/yyyy-mm-dd/*.csv


## Development

To develop `go-twitch-analytics` or interact with its source code in any meaningful way, be
sure you have the following installed:

### Prerequisites

- [Git][git]
- [Go 1.12][golang]+
- [golangCI-Lint][golint]
- [Delve Debugger][dlv]
- [AWS CLI][aws-cli]
- [AWS SAM CLI][aws-sam-cli]

>You will need to activate [Modules][modules] for your version of [GO][golang], 

> by setting the `GO111MODULE=on` environment variable set, or by enabling go-mod in [GoLand][goLand]

### [golangCI-Lint][golint]
```shell
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin latest
```

### [Delve Debugger][dlv]
```shell
GOARCH=amd64 GOOS=linux go build -o ./dlv github.com/go-delve/delve/cmd/dlv
```

### [AWS CLI Configuration][aws-cli-config]
> Make sure you configure the AWS CLI
- AWS Access Key ID
- AWS Secret Access Key
- Default region 'us-east-1'
```shell
aws configure
```

### MISC - [LocalStack][localstack]
<strong>NOTE*:- This repo supports to alteratively use [LocalStack][localstack] 
AWS services locally</strong>

 - S3 
 - Dynamo
 - Lambda
 
Configure the ENV Vars
> note these are the default localstack services host & port
- ```export LOCALSTACK_HOST=localhost```
- ```export LOCALSTACK_DYNAMO=4569```
- ```export LOCALSTACK_S3=4572```
- ```export LOCALSTACK_LAMBDA=4574```

## SETUP
> Steps to deploy serverless stack

- Setup prerequisites
- Create a S3 bucket on AWS to host your packaged serverless stack
- Build the lambda handlers, ensure they're no errors!
- Replace the 'XXXX '**environment** variables in **template.yml** with your credentials!
- Run one of the <strong>package</strong> commands, with the S3 bucket name created above
- Deploy the Serverless stack , with a stack name of your choice!
- Get a user access token with **analytics:read:extensions** OR **analytics:read:games** PERMISISONS 
- Store user access token against your twitchID in the USERS dynamoDB table, example dynamo item can be found below
- Wait until 02:00 / 02:30  (GMT) for the lambdas to be run!
- Observe AWS S3 and AWS Cloudwatch lambda logs

> Example Dynamo Item
```json
{
	"id": "XXXX",
	"user_access_token": {
		"access_token": "XXXXX",
		"expires_in": 0,
		"refresh_token": "XXXXXX",
		"scope": [
			"analytics:read:extensions",
			"analytics:read:games"
		]
	}
}
```
## Build
> How to build the lambda functions

> **Note***:- Executables are placed in 'dist/handlers/*/main' 

```shell
make build
```

**OR**

- **Lint the project FIRST**
> **Note***:-This may throw build **errors**, which **may** not break anything!

```Shell 
golangci-lint run
```
- **Build individual lambda**

> **Note***:- replace '*' with the lambda handler name

```Shell 
GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o dist/*/main  ./lambda/handlers/*/main.go
```

## Running locally with debugging
> How to run the lambda functions locally

> **Note***:- By default debugging is on port:**5859**

### Fetch-Extension-Analytics
> Debug the  Twitch Extension analytics lambda

```shell
sam local invoke -e lambda/test/event.json "FetchExtAnalytics" -d 5859 --debugger-path . --debug-args="-delveAPI=2" --log-file ./output.log
```

### Fetch-Game-Analytics
> Debug the  Twitch game analytics lambda

> **Note***:- please replace '*' with either 'FetchExtAnalytics' or ''
```shell
sam local invoke -e lambda/test/event.json "FetchGameAnalytics" -d 5859 --debugger-path . --debug-args="-delveAPI=2" --log-file ./output.log
```

## Package & Deploy
> This section describes how to validate, compile and deploy the serverless stack.

<strong>NOTE*: Please change the S3 bucket and stack-name within the makefile!!</strong>

### Validate
> Validate the template.yml before packaging!

```shell
make validate
```

**OR**

```shell
sam validate
```

### Package
> How to compile template.yml after updating the serverless template definition.

```shell
make package
```

**OR**

```shell 
sam package --template-file template.yml --s3-bucket {{stack-bucket}} --output-template-file packaged.yml
```

**OR**

```shell 
aws cloudformation package  --template-file template.yml --output-template-file packaged.yml --s3-bucket {{stack-bucket}}
```

### Deploy
> Deploy the serverless stack

```shell
make deploy
```

**OR**

```shell
sam deploy --template-file ./packaged.yml --stack-name {{provide-a-name}} --capabilities CAPABILITY_IAM
```

**OR**

```shell 
aws cloudformation deploy --template-file ./packaged.yml --stack-name {{provide-a-name}} --capabilities CAPABILITY_IAM
```

### Describe Stack
> Get information about the back-end infrastructure

```shell
make describe
```

**OR**

```shell 
aws cloudformation describe-stacks --stack-name {{provide-a-name}} 
```

## Contributors

This project exists thanks to **all** the people who contribute.

## Donations
All donations are appreciated!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](http://paypal.me/crazyjack12)