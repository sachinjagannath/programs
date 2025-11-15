package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/config"
	grpcsrv "github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/internal/grpc"
	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/internal/service"
	"github.com/sachinjangam/programs/golang/food_delivery_app/foodsvc/repo"
)

func main() {
	db := config.ConnectMySql()
	sqlRepo := repo.NewRepo(db)
	svc := service.NewService(sqlRepo)
	server := grpcsrv.NewServer(svc)

	go func() {
		grpcsrv.ServeGrpc(":50051", server)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("shutting down")

}
