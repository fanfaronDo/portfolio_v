package main

import (
	"fmt"
	"github.com/fanfaronDo/portfolio_v/internal/config"
	"github.com/fanfaronDo/portfolio_v/internal/handler"
	"github.com/fanfaronDo/portfolio_v/internal/repository"
	"github.com/fanfaronDo/portfolio_v/internal/server"
	"github.com/fanfaronDo/portfolio_v/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.ConfigLoad()
	server := &server.Server{}
	connectToDB, err := repository.NewMySql(*cfg)
	if err != nil {
		fmt.Println(err)
	}
	repo := repository.NewRepository(connectToDB)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r := handler.InitRoutes()
	go func() {
		if err := server.Run(config.HttpServer.Address, config.HttpServer.Port, r); err != nil {
			fmt.Printf("error occured while running http server: %s\n", string(err.Error()))
			return
		}
	}()

	fmt.Printf("Server started on %s\n", "http://"+config.HttpServer.Address+":"+config.HttpServer.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("Shutting down server...")

	defer server.Stop(nil)

}
