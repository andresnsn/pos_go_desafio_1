package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"pos_go_desafio_1/internal/domain"
)

func NewDb() *gorm.DB {

	// Data source name
	dsn := "file:database.db?cache=shared&mode=rwc&_pragma=journal_mode(WAL)"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	err = db.AutoMigrate(&domain.USDBRL{})
	if err != nil {
		panic("Falha ao executar AutoMigrate: " + err.Error())
	}

	return db
}
