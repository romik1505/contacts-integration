package store

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"os"
	"week3_docker/internal/config"
)

type Store struct {
	DB *sqlx.DB
}

//var embedMigrations embed.FS

func NewStore() (Store, error) {
	dbConf := config.Config.DBConfig
	conString := dbConf.ConnectionString()
	con, err := sqlx.Open(dbConf.Driver, conString)

	if err != nil {
		return Store{}, err
	}

	if err := con.Ping(); err != nil {
		return Store{}, err
	}

	//goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("mysql"); err != nil {
		return Store{}, fmt.Errorf("goose SetDialect: %v", err)
	}

	path, err := os.Getwd()
	if err := goose.Up(con.DB, path+"/migrations"); err != nil {
		return Store{}, fmt.Errorf("goose migrations: %v", err)
	}
	return Store{
		DB: con,
	}, nil
}
func (s Store) Builder() sq.StatementBuilderType {
	return sq.StatementBuilder.RunWith(s.DB)
}
