package DB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) *gorm.DB {
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
