package store

import "database/sql"

func NewNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullInt64(u int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: u,
		Valid: true,
	}
}
