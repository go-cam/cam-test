module test

go 1.14

require (
	github.com/go-cam/cam v0.0.0-20200226085529-b64c3af3eb80
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.23.0
	gorm.io/gorm v0.2.26
)

replace github.com/go-cam/cam => ./../cam
