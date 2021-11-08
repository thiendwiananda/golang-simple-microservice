package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	_logMysqlRepo "github.com/thiendwiananda/golang-simple-microservice/log/repository/mysql"
	_movieGrpcHandler "github.com/thiendwiananda/golang-simple-microservice/movie/delivery/grpc"
	_movieHttpHandler "github.com/thiendwiananda/golang-simple-microservice/movie/delivery/http"
	_movieOmdbRepo "github.com/thiendwiananda/golang-simple-microservice/movie/repository/ombd"
	_movieUsecase "github.com/thiendwiananda/golang-simple-microservice/movie/usecase"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(`env.json`)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	var serverType string
	flag.StringVar(&serverType, "server-type", "http", "choice serving server between rest or grpc")
	flag.Parse()

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	mysqlConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	movieRepo := _movieOmdbRepo.NewOmdbMoviewRepository(viper.GetString(`OMB_key`))
	LogRepo := _logMysqlRepo.NewMysqlLogRepository(mysqlConnection)
	movieUsecase := _movieUsecase.NewMovieUsecase(movieRepo, LogRepo)

	if serverType == "http" {
		server := echo.New()
		_movieHttpHandler.NewMovieHandler(server, movieUsecase)

		log.Fatal(server.Start(viper.GetString(`http_server.port`)))
	}

	list, err := net.Listen("tcp", viper.GetString(`grpc_server.port`))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	_movieGrpcHandler.NewMovieServerGrpc(grpcServer, movieUsecase)
	fmt.Println("GRPC server start on port" + viper.GetString(`grpc_server.port`))

	if err := grpcServer.Serve(list); err != nil {
		log.Fatal(err)
	}
}
