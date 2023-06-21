package model

import "database/sql"

type Account struct {
	ID            sql.NullInt64  `db:"id"`
	Subdomain     sql.NullString `db:"subdomain"`
	AuthCode      sql.NullString `db:"auth_code"`
	IntegrationID sql.NullString `db:"integration_id"`
	AccessToken   sql.NullString `db:"access_token"`
	RefreshToken  sql.NullString `db:"refresh_token"`
	Expires       sql.NullInt64  `db:"expires"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

type ListAccountFilter struct {
	Page        int
	Limit       int
	NeedRefresh bool
}
