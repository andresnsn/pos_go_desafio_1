package database

import (
	"context"
	"pos_go_desafio_1/internal/domain"

	"gorm.io/gorm"
)

type USDBRLRepository struct {
	Db *gorm.DB
}

func (t *USDBRLRepository) Save(ctx context.Context, usdbrl *domain.USDBRL) error {
	tx := t.Db.WithContext(ctx).Save(usdbrl)
	return tx.Error
}
