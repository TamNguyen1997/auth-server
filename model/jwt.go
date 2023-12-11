package model

type BasicJwt struct {
	Authorities []string `json:"authorities"`
	Exp         int64    `json:"exp"`
}

func (jwt *BasicJwt) Valid() error {
	return nil
}
