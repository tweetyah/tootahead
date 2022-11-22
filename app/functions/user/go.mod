module user

go 1.19

replace core => ../core

require (
	github.com/aws/aws-lambda-go v1.34.1
	github.com/bmorrisondev/go-utils v1.0.1
)

require (
	github.com/aws/aws-sdk-go v1.44.131 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)
