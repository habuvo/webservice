package storage

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/habuvo/webservice/pkg/loggers"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type PostgreStorage struct {
	connPool *pgxpool.Pool
}

func NewPostgreStorage(addr string, log *zap.Logger) (*PostgreStorage, error) {
	cfg, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, err
	}
	cfg.ConnConfig.Logger = loggers.NewZapLogger(log)
	connPool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	conn, err := connPool.Acquire(context.Background())
	if err != nil {
		connPool.Close()
		return nil, err
	}

	conn.Release()

	return &PostgreStorage{connPool: connPool}, nil
}

func migrateUp(cn *pgxpool.Conn, schema string) (err error) {
	_, err = cn.Exec(context.Background(), schema)
	return
}
