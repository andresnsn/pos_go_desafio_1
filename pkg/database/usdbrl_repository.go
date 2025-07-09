package database

import (
	"context"
	"errors"
	"fmt"
	"pos_go_desafio_1/internal/domain"
	"time"

	"gorm.io/gorm"
)

type USDBRLRepository struct {
	Db *gorm.DB
}

func (r *USDBRLRepository) Save(usdbrl *domain.USDBRL) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	if err := r.Db.WithContext(ctx).Save(usdbrl).Error; err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("O timeout de 10ms para salvar no banco de dados foi excedido.")
		}
	}

	tx := r.Db.Save(usdbrl)
	return tx.Error
}
