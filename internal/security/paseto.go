package security

import (
	"errors"
	"time"

	"github.com/o1egl/paseto"
)

type PasetoGen struct {
	paseto      *paseto.V2
	symetricKey []byte
}

func NewMockPasetoGen() *PasetoGen {
	return &PasetoGen{
		paseto:      paseto.NewV2(),
		symetricKey: []byte("1234 1234 1234 1234 1234 1234 12"),
	}
}

func NewPasetoGen(symetricKey []byte) *PasetoGen {
	return &PasetoGen{
		paseto:      paseto.NewV2(),
		symetricKey: symetricKey,
	}
}

func (p *PasetoGen) CreateToken(username string) (string, error) {
	payload := NewPayload(username, 2*time.Hour)
	return p.paseto.Encrypt(p.symetricKey, payload, nil)
}

func (p *PasetoGen) VerifyToken(token string) (*Payload, error) {
	payload := new(Payload)
	err := p.paseto.Decrypt(token, p.symetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	if payload.IsExpired() {
		return nil, errors.New("payload is expired")
	}

	return payload, nil
}
