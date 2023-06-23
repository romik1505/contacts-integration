package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAccounts, downAccounts)
}

func upAccounts(tx *sql.Tx) error {
	_, err := tx.Exec(
		"create table accounts" +
			"(" +
			"id SERIAL," +
			"subdomain VARCHAR(255) NOT NULL, " +
			"auth_code TEXT," +
			"integration_id VARCHAR(255)," +
			"access_token TEXT," +
			"refresh_token TEXT," +
			"expires BIGINT UNSIGNED," +
			"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
			"updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" +
			");")
	if err != nil {
		return err
	}
	return nil
}

func downAccounts(tx *sql.Tx) error {
	_, err := tx.Exec("drop table accounts;")
	if err != nil {
		return err
	}
	return nil
}
