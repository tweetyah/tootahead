module tester

go 1.19

replace github.com/tweetyah/lib => ../lib

require (
	github.com/joho/godotenv v1.4.0
	github.com/tweetyah/lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-lambda-go v1.34.1 // indirect
	github.com/aws/aws-sdk-go v1.43.3 // indirect
	github.com/bmorrisondev/go-utils v1.0.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)
