package api

import (
	"time"
)

const (
	HANDLERBASEGROUP = "/api/v1"
)

type TestStruct struct {
	Id int
	Username string
	Password string
	Email string
	CreatedOn time.Time
}