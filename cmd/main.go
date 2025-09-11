package main

import (
	"os"
	"todo-list/pkg/handler"
	"todo-list/pkg/repository"
	"todo-list/pkg/repository/users"
	"todo-list/pkg/services"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed initializing db: %s", err.Error())
	}
	repos := users.NewRepository(db)
	services := services.NewService(repos)
	handlers := handler.NewHandler(services.Authorization, services.TodoList)
	routes := handlers.InitRoutes()

	if err := routes.Run(":" + viper.GetString("port")); err != nil {
		logrus.Fatalf("gin run error: %v", err)
	}

}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
