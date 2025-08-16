package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
	"withdrawal-balance/internal/api/xendit"
	"withdrawal-balance/internal/domain/user"
	"withdrawal-balance/internal/domain/wallet"
	"withdrawal-balance/internal/domain/withdrawalhistory"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"withdrawal-balance/internal/rest"
	"withdrawal-balance/internal/rest/middleware"

	"github.com/joho/godotenv"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//prepare database
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()
	// prepare echo

	e := echo.New()
	e.Use(middleware.CORS)
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Prepare Repository
	userRepo := user.NewRepository(dbConn)
	walletRepo := wallet.NewRepository(dbConn)
	withdrawalHistoryRepo := withdrawalhistory.NewRepository(dbConn)
	xenditRepo := xendit.NewRepository()

	// Build service Layer
	userSvc := user.NewService(userRepo)
	walletSvc := wallet.NewService(walletRepo)
	xenditSvc := xendit.NewService(xenditRepo)
	withdrawalHistorySvc := withdrawalhistory.NewService(withdrawalHistoryRepo, userSvc, walletSvc, xenditSvc)

	// Build rest Layer
	rest.NewWithdrawalHistoryHandler(e, withdrawalHistorySvc)
	rest.NewWalletHandler(e, walletSvc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}
