package db

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    db *gorm.DB
    postgresConnection string
)

func Init() {
    var err error

    postgresConnection = fmt.Sprintf("host=%s port=%s user=%s dbname=%s " +
        "sslmode=%s password=%s",
        "192.168.99.100", "32770", "root", "blog_api", "disable", "root")

    db , err = gorm.Open("postgres", postgresConnection)

    if err != nil {
        panic(err)
        return
    }
}

func GetDB() *gorm.DB {
    return db;
}
