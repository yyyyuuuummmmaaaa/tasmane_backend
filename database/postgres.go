package database


import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "fmt"
)

func Connect() *gorm.DB {

    dsn := "host=postgres user=root password=gopass dbname=godb port=5432 sslmode=disable TimeZone=Asia/Tokyo"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Println(err)
    }

    return db
}

