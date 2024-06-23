package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
	web_go "web-go"
	"web-go/Model"
	"web-go/pkg/handler"
	"web-go/pkg/repo"
	"web-go/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("init config fail", err.Error())
	}

	db, err := gorm.Open("postgres", viper.GetString("database.url"))
	if err != nil {
		log.Fatal("init db fail", err.Error())
	} else {
		log.Println("init db success")
	}
	defer db.Close()

	if !db.HasTable(&Model.User{}) {
		if err := db.AutoMigrate(&Model.User{}).Error; err != nil {
			log.Fatalf("failed to create table: %v", err)
		}
	}

	if !db.HasTable(&Model.TodoItem{}) {
		if err := db.AutoMigrate(&Model.TodoItem{}).Error; err != nil {
			log.Fatalf("failed to create table: %v", err)
		}
	}

	repos := repo.NewRepo(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(web_go.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("Config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
