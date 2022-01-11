package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbConfig struct {
	dbHost, dbName, dbUser, dbPassword, mongoHost string
	dbPort                                        uint
}

func main() {

	dbConfig, err := testDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConnection, err := ConnectSQL("mysql", dbConfig.dbHost, dbConfig.dbUser, dbConfig.dbPassword, dbConfig.dbName, dbConfig.dbPort)

	if err != nil {
		log.Fatal("Could not connect to database", err)
		return
	}

	sqlDB, err := dbConnection.DB()
	if err != nil {
		log.Fatal("No connection in gorm DB()", err)
		return
	}

	defer sqlDB.Close()

	if err = MigrateTables(dbConnection); err != nil {
		log.Fatal(err)
		return
	}
	router := mux.NewRouter()

	SetupRoutes(dbConnection, router)

	serverStopChan, subsShutdownChan := CreateServerStopChan()

	<-serverStopChan
	close(serverStopChan)
	subsShutdownChan <- true

}

func CreateServerStopChan() (chan os.Signal, chan bool) {
	// serverStopChan to stop the server based on the os signels. like os Interrupt
	serverStopChan := make(chan os.Signal, 1)
	subsShutdownChan := make(chan bool, 1)
	// passing channel to os signel notify
	signal.Notify(serverStopChan, os.Interrupt, syscall.SIGTERM)
	return serverStopChan, subsShutdownChan
}

func testDBConfig() (dbConfig, error) {
	config := dbConfig{}

	config.dbName = "ArticleManagement"
	config.dbPassword = "password"
	config.dbUser = "RM"
	config.dbHost = "newdb"
	config.dbPort = 3306

	return config, nil
}

func MigrateTables(dbConnection *gorm.DB) error {

	log.Println("Starting migration")
	if err := migrateAll(dbConnection, &Article{}); err != nil {
		return err
	}

	return nil
}

func migrateAll(dbConnection *gorm.DB, tables ...interface{}) error {

	for _, model := range tables {

		modelName := getType(model)

		if err := dbConnection.AutoMigrate(model); err != nil {

			return errors.Wrapf(err, "Could not migrate %s", modelName)
		}

	}

	return nil
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

var dbConnection *gorm.DB

func ConnectSQL(dialect, dbHost, dbUser, dbPass, dbName string, dbPort uint) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             40 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	gormConf := &gorm.Config{
		Logger: newLogger,
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=false&autocommit=true&charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	dialector := mysql.Open(connString)
	conn, err := gorm.Open(dialector, gormConf)

	if err != nil {
		return nil, err
	}

	dbConnection = conn

	return dbConnection, nil
}
