// +heroku goVersion go1.14.4
// +heroku install ./cmd/...
module github.com/vivaldy22/eatnfit-food-service

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	github.com/spf13/viper v1.7.1
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.24.0
)
