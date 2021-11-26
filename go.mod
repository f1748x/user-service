module user-service

go 1.16

require (
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20211124092129-2acede53f3e8
	github.com/go-kratos/kratos/v2 v2.1.2
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/api v1.11.0
	google.golang.org/genproto v0.0.0-20211016002631-37fc39342514
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.3
)
