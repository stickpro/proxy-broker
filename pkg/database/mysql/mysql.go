package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDataBase(host, user, password, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	fmt.Println(dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
