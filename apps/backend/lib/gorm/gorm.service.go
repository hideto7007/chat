package lib

import (
	"chat/models"
	"log"

	"gorm.io/gorm"
)


// 開発用：テーブルを初期化（全データ削除）
func DropAndCreateTables(db *gorm.DB) {
	for _, model := range models.AllModels() {
		if err := db.Migrator().DropTable(model); err != nil {
			log.Fatalf("failed to drop table: %v", err)
		}
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate: %v", err)
		}
	}
}

// 本番・開発共通：マイグレーションのみ
func Migrate(db *gorm.DB) {
	for _, model := range models.AllModels() {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate: %v", err)
		}
	}
}