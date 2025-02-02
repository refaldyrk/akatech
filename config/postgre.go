package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgres() *gorm.DB {
	fmt.Println("Initial Postgre Start")
	db, err := gorm.Open(postgres.Open(viper.GetString("POSTGRE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Initial Postgre Done")

	return db
}
