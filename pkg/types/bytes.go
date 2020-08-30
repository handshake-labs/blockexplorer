package types

import (
	"encoding/hex"
	"encoding/json"
)

type Bytes []byte

func (b Bytes) MarshalText() ([]byte, error) {
	s := hex.EncodeToString([]byte(b))
	return []byte(s), nil
}

func (b *Bytes) UnmarshalText(t []byte) error {
	h, err := hex.DecodeString(string(t))
	*b = Bytes(h)
	return err
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	s := hex.EncodeToString([]byte(b))
	return json.Marshal(s)
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return b.UnmarshalText([]byte(s))
}
