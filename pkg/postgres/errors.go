package postgres

import (
	"errors"
)

var ErrTxClosed = errors.New("tx closed")
