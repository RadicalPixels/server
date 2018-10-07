package util

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// ParseContinuousColorHexString ...
func ParseContinuousColorHexString(colorStr string) []string {
	if colorStr == "" {
		return nil
	}
	slice := []byte(colorStr)
	const colorlen = 6
	var colors []string
	size := len(slice)

	for i := 0; i < size; i += colorlen {
		if i%colorlen == 0 {
			x := (i + colorlen)
			if x > size {
				break
			}
			colors = append(colors, "#"+string(slice[i:x]))
		}
	}

	return colors
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
