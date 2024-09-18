package main

import (
	"github.com/fanfaronDo/portfolio_v/internal/config"
	"github.com/fanfaronDo/portfolio_v/internal/repository"
	"github.com/fanfaronDo/portfolio_v/internal/service"
	"log"
)

func main() {
	cfg := config.ConfigLoad()
	conn, err := repository.NewMySql(*cfg)
	if err != nil {
		log.Println("panic occurred:", err)
	}
	repo := repository.NewRepository(conn)
	service := service.NewService(repo)

	service.Delete(3)
	//p := handler.NewTemplatePaginationPage(20, 7)
	//fmt.Println(p)
	//d, err := service.Projects.GetRange(p.StartProjectEdge, p.EndProjectEdge)
	//fmt.Println(d, err)
}
