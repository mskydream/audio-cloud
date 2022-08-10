package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/audio-cloud/config"
)

// type DB struct {
// 	Conn *sqlx.DB
// }

func InitDatabase(c *config.DB) *sqlx.DB {
	source := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, c.Password, c.Address, c.Name)

	db, err := sqlx.Connect("pgx", source)
	if err != nil {
		panic(err)
	}

	m, err := migrate.New("file:migrations", source)
	if err != nil {
		panic(err)
	}
	if err = m.Up(); err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}

	return db
}
