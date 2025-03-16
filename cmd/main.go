package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	urlshortener "github.com/ramil66/url-shortener"
	"github.com/ramil66/url-shortener/pkg/handler"
	"github.com/ramil66/url-shortener/pkg/repository"
	"github.com/ramil66/url-shortener/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitRoutesWithCORS(handlers *handler.Handler) *gin.Engine {
	router := gin.New()

	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешить любые источники
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Подключение маршрутов
	api := router.Group("/api")
	{
		handlers.InitRoutes(api)
	}

	return router
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error run config: %s", err.Error())
	}

	db, err := repository.NewPostregDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error run database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// server := new(urlshortener.Server)
	// if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	// 	logrus.Fatalf("error run server: %s", err.Error())
	// }
	router := InitRoutesWithCORS(handlers)

	server := new(urlshortener.Server)
	if err := server.Run(viper.GetString("port"), router); err != nil {
		logrus.Fatalf("error run server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
