package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

// PostgresDB ...
type PostgresDB struct {
	*pgxpool.Pool
	Sq *Squirrel
}

// New provides PostgresDB struct init
func New(
	host,
	port,
	name,
	user,
	password,
	sslmode string,
) (*PostgresDB, error) {
	db := PostgresDB{Sq: NewSquirrel()}
	connectConfig := configToStr(host,
		port,
		name,
		user,
		password,
		sslmode,
	)
	if err := db.connect(connectConfig); err != nil {
		return nil, err
	}
	return &db, nil
}

func (p *PostgresDB) connect(connectConfig string) error {
	pgxpoolConfig, err := pgxpool.ParseConfig(connectConfig)
	if err != nil {
		return fmt.Errorf("unable to parse database config: %w", err)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), pgxpoolConfig)
	if err != nil {
		return fmt.Errorf("unable to connect database config: %w", err)
	}
	p.Pool = pool
	return nil
}

func (p *PostgresDB) Close() {
	p.Pool.Close()
}

func configToStr(
	host,
	port,
	name,
	user,
	password,
	sslmode string,
) string {
	var conn []string
	if len(host) != 0 {
		conn = append(conn, "host="+host)
	}
	if len(port) != 0 {
		conn = append(conn, "port="+port)
	}
	if len(user) != 0 {
		conn = append(conn, "user="+user)
	}
	if len(password) != 0 {
		conn = append(conn, "password="+password)
	}
	if len(name) != 0 {
		conn = append(conn, "dbname="+name)
	}
	if len(sslmode) != 0 {
		conn = append(conn, "sslmode="+sslmode)
	}
	return strings.Join(conn, " ")
}
