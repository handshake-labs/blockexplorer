package types

import (
	"encoding/json"
	"fmt"
  "errors"
  "github.com/randomlogin/decimal"
)

const _MONEY_SCALE float64 = 1e6

type Money int64

func (m Money) MarshalJSON() ([]byte, error) {
	i := int64(m)
	f := float64(m) / _MONEY_SCALE
	if int64(f*_MONEY_SCALE) != i {
		return nil, fmt.Errorf("json: unsupported value for Money marshal: %i", i)
	}
	return json.Marshal(float64(m) / _MONEY_SCALE)
}

func (m *Money) UnmarshalJSON(data []byte) error {
  var floatValue decimal.Decimal
	err := floatValue.UnmarshalJSON(data)
	if err != nil {
		return err
	}
  i := floatValue.Mul(decimal.NewFromFloat(_MONEY_SCALE))
  if !i.IsInteger() {
    return errors.New("json: Could not make int value for money")
  }
  *m = Money(i.IntPart())
  return nil
}

