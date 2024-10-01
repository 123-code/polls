package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	//"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	encodedPassword := url.QueryEscape("AVNS_LNgImquHJXNIMn4aMTt")

	dsn := fmt.Sprintf("postgres://avnadmin:AVNS_Xv4rbmtFkoxKOGBDks_@pg-382a4bdb-udla-54df.g.aivencloud.com:18022/defaultdb?sslmode=require", encodedPassword)
	//dsn := "postgres://avnadmin:AVNS_LNgImquHJXNIMn4aMTt@actixwebpostgres-udla-54df.aivencloud.com:18022/defaultdb?sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})


	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Connected to database")
	}
}
