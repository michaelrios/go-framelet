package repositories

import "golang.org/x/xerrors"

var (
	ErrEntityNotFound = xerrors.New("entity not found")
	ErrResourceFailed = xerrors.New("DB hung up")
	)

