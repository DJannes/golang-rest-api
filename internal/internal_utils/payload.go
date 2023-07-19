package internal_utils

import "time"

type Payload struct {
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, tokenTimeValidity time.Duration) *Payload {
	return &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(tokenTimeValidity),
	}
}

func (p *Payload) IsExpired() bool {
	return time.Now().After(p.ExpiredAt)
}
