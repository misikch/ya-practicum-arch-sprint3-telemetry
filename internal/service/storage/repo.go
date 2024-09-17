package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db     *sqlx.DB
	logger Log
}

func New(db *sqlx.DB, logger Log) *Storage {
	return &Storage{
		db:     db,
		logger: logger,
	}
}

func (s *Storage) SetDeviceStatus(ctx context.Context, deviceUUID string, status string, author string) (err error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			rollbackErr := tx.Rollback()
			if err != nil {
				s.logger.Error("recovered: failed to rollback transaction", rollbackErr)
			}
		} else if err != nil {
			rollbackErr := tx.Rollback()
			if err != nil {
				s.logger.Error("failed to rollback transaction", rollbackErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	// Обновление статуса устройства
	_, err = tx.ExecContext(ctx, "UPDATE devices SET status = ?, updated_at = NOW() WHERE uuid = ?", status, deviceUUID)
	if err != nil {
		return err
	}

	// Логирование изменения статуса
	_, err = tx.ExecContext(ctx,
		"INSERT INTO device_status_log (device_uuid, status, author, changed_at) VALUES (?, ?, ?, NOW())",
		deviceUUID, status, author)
	if err != nil {
		return
	}

	return
}
