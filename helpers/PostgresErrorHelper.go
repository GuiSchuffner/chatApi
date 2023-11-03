package helper

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const DUPLICATED_KEY = "23505"

func IsDuplicatedKeyError(err error) bool {
	var perr *pgconn.PgError
	if errors.As(err, &perr) {
		return perr.Code == DUPLICATED_KEY
	}
	return false
}
