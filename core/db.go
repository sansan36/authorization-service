package core

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	userv1 "github.com/sansan36/authorization-service/gen/user/v1"
	"github.com/sansan36/authorization-service/libs"
	"github.com/sansan36/authorization-service/libs/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

var (
	DBMain    *gorm.DB
	DBMainSQL *sql.DB
	logLevel  gormLog.LogLevel = gormLog.Info

	POSTGRES_HOST     = libs.GetEnv("POSTGRES_HOST", "")
	POSTGRES_PORT     = libs.GetEnv("POSTGRES_PORT", "")
	POSTGRES_USER     = libs.GetEnv("POSTGRES_USER", "")
	POSTGRES_PASSWORD = libs.GetEnv("POSTGRES_PASSWORD", "")
	POSTGRES_DB       = libs.GetEnv("POSTGRES_DB", "")

	OrmList []interface{} = []interface{}{
		&userv1.UserORM{},
	}
)

func StartDBConnection() {
	logger.Printf("Starting Db Connections...")

	InitDBMain()
}

func InitDBMain() {
	logger.Printf("Main Db - Connecting")
	var err error

	gormLogger := gormLog.New(
		logger.Logger,
		gormLog.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		POSTGRES_HOST,
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_DB,
		POSTGRES_PORT,
	)
	fmt.Println(dsn)
	DBMain, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logger.Fatalf("Failed connect to DB main: %v", err)
		os.Exit(1)
		return
	}

	DBMainSQL, err = DBMain.DB()
	if err != nil {
		logger.Fatalf("Error cannot initiate connection to DB main: %v", err)
		os.Exit(1)
		return
	}

	DBMainSQL.SetMaxIdleConns(0)
	DBMainSQL.SetMaxOpenConns(0)

	err = DBMainSQL.Ping()
	if err != nil {
		logger.Fatalf("Cannot ping DB main: %v", err)
		os.Exit(1)
		return
	}

	logger.Printf("Main Db - Connected")
}

func CloseDBMain() {
	logger.Print("Closing DB Main Connection ... ")
	if err := DBMainSQL.Close(); err != nil {
		logger.Fatalf("Error on disconnection with DB Main : %v", err)
	}
	logger.Println("Closing DB Main Success")
}

func MigrateDB() error {
	InitDBMain()
	defer CloseDBMain()

	logger.Println("Migration process begin...")
	if err := DBMain.AutoMigrate(OrmList...); err != nil {
		logger.Fatalf("Migration failed: %v", err)
		os.Exit(1)
	}

	logger.Println("Migration process finished...")

	return nil
}
