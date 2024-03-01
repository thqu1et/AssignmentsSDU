package SQL

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {
	var err error
	err = godotenv.Load(".env")

	var user = os.Getenv("MYSQL_USER")
	var password = os.Getenv("MYSQL_PASSWORD")
	var net = os.Getenv("MYSQL_NET")
	var address = os.Getenv("MYSQL_Address")
	var DBName = os.Getenv("MYSQL_Name")

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, net, address, DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
