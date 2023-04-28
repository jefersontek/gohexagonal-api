package util

import "github.com/ericlagergren/decimal"

func FloatToDecimal(value float64) *decimal.Big {
	dec := new(decimal.Big)
	return dec.SetFloat64(value).Quantize(2)
}

// DecimalToFloat
func DecimalToFloat(cost *decimal.Big) float64 {
	var costFloat float64
	if cost == nil {
		return costFloat
	}
	if value, ok := cost.Quantize(2).Float64(); ok {
		costFloat = value
	}
	return costFloat
}
