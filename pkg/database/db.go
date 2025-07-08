package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"pos_go_desafio_1/internal/domain"
)

func NewDb() *gorm.DB {

	// Data source name
	dsn := "file:database.db?cache=shared&mode=rwc&_pragma=journal_mode(WAL)"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Falha ao obter *sql.DB: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)

	err = db.AutoMigrate(&domain.USDBRL{})
	if err != nil {
		panic("Falha ao executar AutoMigrate: " + err.Error())
	}

	return db
}
