package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"latentlab.cc/easyhousing/pkg/model"
	"os"
	"time"
)

// CreateClient Creates a new gorm client with the database connection using environment variables.
func CreateClient() (*gorm.DB, error) {
	cfg := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "app", os.Getenv("PORT")),
	}

	psql := postgres.New(cfg)

	gormCfg := gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	gormDB, err := gorm.Open(psql, &gormCfg)
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

func SetupTables() error {
	gormDB, err := CreateClient()
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&model.Home{})
	if err != nil {
		return err
	}
	err = gormDB.AutoMigrate(&model.SupportingDoc{})
	if err != nil {
		return err
	}
	err = gormDB.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
