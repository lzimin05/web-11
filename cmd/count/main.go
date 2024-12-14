package main

import (
	"flag"
	"lab-11/internal/count/api"
	"lab-11/internal/count/config"
	"lab-11/internal/count/provider"
	"lab-11/internal/count/usecase"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "./configs/hello_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(cfg.Usecase.DefaultMessageCount, prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
