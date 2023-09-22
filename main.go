package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"transaction/handler"
	"transaction/repository"
	"transaction/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	bankRepositoryDB := repository.NewBankRepositoryDB(db)
	bankService := service.NewBankService(bankRepositoryDB)
	bankHandler := handler.NewBankHandler(bankService)

	transRepositoryDB := repository.NewTransactionRepositoryDB(db)
	transService := service.NewTransactionService(transRepositoryDB)
	transHandler := handler.NewTransactionHandler(transService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", customerHandler.PostCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers", customerHandler.PutCustomer).Methods(http.MethodPut)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.DeleteCustomer).Methods(http.MethodDelete)

	router.HandleFunc("/banks", bankHandler.GetBanks).Methods(http.MethodGet)
	router.HandleFunc("/banks/{bank_id:[0-9]+}", bankHandler.GetBank).Methods(http.MethodGet)
	router.HandleFunc("/banks", bankHandler.PostBank).Methods(http.MethodPost)
	router.HandleFunc("/banks", bankHandler.PutBank).Methods(http.MethodPut)
	router.HandleFunc("/banks/{bank_id:[0-9]+}", bankHandler.DeleteBank).Methods(http.MethodDelete)

	router.HandleFunc("/transactions", transHandler.GetTransactions).Methods(http.MethodGet)
	router.HandleFunc("/transactions/{customer_id:[0-9]+}/customer", transHandler.GetTransactionByCustomerId).Methods(http.MethodGet)
	router.HandleFunc("/transactions/{customer_id:[0-9]+}/bank", transHandler.GetTransactionByBankId).Methods(http.MethodGet)
	router.HandleFunc("/transactions", transHandler.PostTransaction).Methods(http.MethodPost)
	router.HandleFunc("/transactions", transHandler.PutTransaction).Methods(http.MethodPut)
	router.HandleFunc("/transactions/{transactions_id:[0-9]+}", transHandler.DeleteTransaction).Methods(http.MethodDelete)

	log.Printf("go lang started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
