package main

import (
	"os"

	"github.com/EusRique/codepix/application/grpc"
	"github.com/EusRique/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
