package models

import (
	"chat/lib/env"
	"flag"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB{
    var dbHost string
    flag.StringVar(&dbHost, "dbhost", "", "Database host")
    flag.Parse()

    // 引数が渡された場合のみ有効化
    if dbHost == "" {
        dbHost = env.GetEnv("DATABASE_URL", "db")
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        dbHost,
        env.GetEnv("DB_USER", "local"),
        env.GetEnv("DB_PASSWORD", "password"),
        env.GetEnv("DB_NAME", "postgres"),
        env.GetEnv("DB_PORT", "5432"),
        env.GetEnv("DB_SSLMODE", "disable"),
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    log.Println("dsn", dsn)
    return db
}

func AllModels() []interface{} {
    return []interface{}{
        &User{},
        // &OtherModel{}, // 追加していく
    }
}
