package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Storage struct {
	client *mongo.Client
	logger *zap.SugaredLogger
	dbName string
}

// New создаёт новый экземпляр хранилища (Storage) с использованием MongoDB.
func New(client *mongo.Client, logger *zap.SugaredLogger, dbName string) *Storage {
	return &Storage{
		client: client,
		logger: logger,
		dbName: dbName,
	}
}

// Database предоставляет доступ к базе данных MongoDB.
func (s *Storage) Database() *mongo.Database {
	return s.client.Database(s.dbName)
}
