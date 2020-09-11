package config

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/vivaldy22/eatnfit-food-service/master/packet"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vivaldy22/eatnfit-food-service/master/food"
	food_service "github.com/vivaldy22/eatnfit-food-service/proto"
	"github.com/vivaldy22/eatnfit-food-service/tools/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitDB() (*sql.DB, error) {
	dbUser := viper.ViperGetEnv("DB_USER", "root")
	dbPass := viper.ViperGetEnv("DB_PASSWORD", "password")
	dbHost := viper.ViperGetEnv("DB_HOST", "localhost")
	dbPort := viper.ViperGetEnv("DB_PORT", "3306")
	schemaName := viper.ViperGetEnv("DB_SCHEMA", "schema")
	driverName := viper.ViperGetEnv("DB_DRIVER", "mysql")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, schemaName)
	dbConn, err := sql.Open(driverName, dbPath)

	if err != nil {
		return nil, err
	}

	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	return dbConn, nil
}

func RunServer(db *sql.DB) {
	host := viper.ViperGetEnv("GRPC_FOOD_HOST", "localhost")
	port := viper.ViperGetEnv("GRPC_FOOD_PORT", "1011")

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	foodService := food.NewService(db)
	food_service.RegisterFoodCRUDServer(srv, foodService)
	packetService := packet.NewService(db)
	food_service.RegisterPacketCRUDServer(srv, packetService)

	reflection.Register(srv)

	log.Printf("Starting GRPC Eat N' Fit Food Server at %v port: %v\n", host, port)
	if err = srv.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
