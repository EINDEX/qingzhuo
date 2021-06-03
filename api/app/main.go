package main

import (
	"fmt"
	"github.com/eindex/qing-zhuo/api/domain"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"

	_postDeliveryHttp "github.com/eindex/qing-zhuo/api/post/delivery/http"
	_middleware "github.com/eindex/qing-zhuo/api/post/delivery/http/middleware"
	_postRepo "github.com/eindex/qing-zhuo/api/post/repository/mysql"
	_postUsecase "github.com/eindex/qing-zhuo/api/post/usecase"
)

func main() {

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file %s \n", err))
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.database"),
	)
	fmt.Printf(dsn)
	DBConn, _ := gorm.Open(mysql.Open(dsn))
	err := DBConn.AutoMigrate(&domain.Post{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router := gin.Default()
	timeoutContext := time.Duration(viper.GetInt("server.timeout")) * time.Second

	pr := _postRepo.NewMysqlPostRepository(DBConn)
	api := router.Group("api", _middleware.PermissionApplyMiddleware())
	pu := _postUsecase.NewPostUsecase(pr, timeoutContext)

	_postDeliveryHttp.NewPostHandler(api.Group("posts"), pu)
	_postDeliveryHttp.NewArchiveHandler(api.Group("archives"), pu)

	router.Run()
}
