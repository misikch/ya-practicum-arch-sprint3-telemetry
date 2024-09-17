package container

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"device-manager/internal/config"
	"device-manager/internal/service/storage"
)

type Container struct {
	Database *sqlx.DB
	Logger   *zap.SugaredLogger
	Storage  *storage.Storage
}

func InitContainer(cfg *config.Config) (*Container, error) {
	c := Container{}
	err := c.initDB(cfg)
	if err != nil {
		return nil, err
	}
	err = c.InitLogger()
	if err != nil {
		return nil, err
	}

	c.initStorage()

	return &c, nil
}

func (c *Container) initDB(cfg *config.Config) error {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		cfg.PgUser,
		cfg.PgPass,
		cfg.PgHost,
		cfg.PgPort,
		cfg.PgDatabase,
	)

	db, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to pg database: %w", err)
	}

	c.Database = db

	return nil
}

func (c *Container) InitLogger() error {
	logger, err := zap.NewDevelopment() //TODO: выбирать окружение
	if err != nil {
		return err
	}

	c.Logger = logger.Sugar()

	return nil
}

func (c *Container) initStorage() {
	c.Storage = storage.New(c.Database, c.Logger)
}
