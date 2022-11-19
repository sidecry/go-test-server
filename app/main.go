package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_articleHttpDelivery "github.com/grimoh/go-test-server/article/delivery/http"
	_articleHTTPDeliveryMiddleware "github.com/grimoh/go-test-server/article/delivery/http/middleware"
	_articleRepo "github.com/grimoh/go-test-server/article/repository/mysql"
	_articleUsecase "github.com/grimoh/go-test-server/article/usecase"
	_authorRepo "github.com/grimoh/go-test-server/author/repository/mysql"
)

func init() {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Japan")

	dsn := fmt.Sprintf("%s?%s", conn, val.Encode())

	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middle := _articleHTTPDeliveryMiddleware.InitMiddleware()
	e.Use(middle.CORS)

	authorRepo := _authorRepo.NewMySQLAuthorRepository(dbConn)
	ar := _articleRepo.NewMySQLArticleRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _articleUsecase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	_articleHttpDelivery.NewArticleHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
