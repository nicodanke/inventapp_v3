package token

import "time"

type Maker interface {
	CreateToken(userId int64, accountId int64, accountCode string, duration time.Duration) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
