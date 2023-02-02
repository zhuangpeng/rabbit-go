package dbx

import (
	"database/sql"
	"time"
)

func NewNullString(s string) sql.NullString {
    if len(s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
        String: s,
        Valid:  true,
    }
}

func NewNullTime(s *time.Time) sql.NullTime {
    if s == nil {
        return sql.NullTime{}
    }
    return sql.NullTime{
		Time:  *s,
		Valid: true,
	}
}