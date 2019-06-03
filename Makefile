clean:
	@echo Cleaning dist folder
	@rm -rf dist
	@mkdir -p dist

build: clean
	@echo running lint
	- golangci-lint run
	@echo building lambda handlers
	@for dir in `ls lambda/handlers/`; do \
		GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o dist/$$dir/main lambda/handlers/$$dir/main.go; \
	done

run-api:
	sam local start-api  --skip-pull-image -d 5859 --debugger-path . --debug-args="-delveAPI=2" --log-file ./output.log

validate:
	sam validate

package: validate build
		sam package \
			--template-file template.yml \
			--s3-bucket {{stack-bucket}} \
			--output-template-file packaged.yml \

deploy:
	sam deploy \
	--template-file ./packaged.yml \
	--stack-name {{provide-a-name}} \
	--capabilities CAPABILITY_IAM \

describe:
		@aws cloudformation describe-stacks \
			--region us-east-1 \
			--stack-name {{provide-a-name}} \
