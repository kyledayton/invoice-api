package invoice

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"invoice-api/numberfmt"
)

type Price struct {
	amountCents uint64
}

func NewPrice(dollars, cents uint64) Price {
	amountCents := (dollars * 100) + cents
	return Price{amountCents}
}

func NewPriceFromFloat(price float64) (Price, error) {
	if price < 0 {
		return Price{0}, errors.New("price cannot be negative")
	}

	amountCents := math.RoundToEven(price * 100)
	amountCents = math.Abs(amountCents)

	return Price{uint64(amountCents)}, nil
}

func NewPriceFromString(str string) (Price, error) {
	str = strings.ReplaceAll(str, "$", "")
	str = strings.ReplaceAll(str, ",", "")

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return Price{0}, err
	}

	return NewPriceFromFloat(val)
}

func (p Price) Dollars() uint64 {
	return p.amountCents / 100
}

func (p Price) Cents() uint64 {
	return p.amountCents % 100
}

func (p Price) Add(rhs Price) Price {
	return Price{p.amountCents + rhs.amountCents}
}

func (p Price) Mul(rhs float64) Price {
	amountCents := float64(p.amountCents) * rhs
	amountCents = math.RoundToEven(amountCents)

	return Price{uint64(amountCents)}
}

func (p Price) String() string {
	dollars := p.Dollars()
	cents := p.Cents()

	return fmt.Sprintf("$%s.%02d", numberfmt.ThousandsSeparated(dollars, ","), cents)
}

func (p *Price) UnmarshalJSON(data []byte) error {
	var raw interface{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	stringVal, isString := raw.(string)
	if isString {
		price, err := NewPriceFromString(stringVal)
		if err != nil {
			return err
		}

		p.amountCents = price.amountCents
		return nil
	}

	numberVal, isNumber := raw.(float64)
	if isNumber {
		price, err := NewPriceFromFloat(numberVal)
		if err != nil {
			return err
		}

		p.amountCents = price.amountCents
		return nil
	}

	arrayVal, isArray := raw.([]interface{})
	if isArray {
		if len(arrayVal) > 2 {
			return errors.New("price array has too many elements")
		}

		floatVals := make([]float64, 2)
		for idx, v := range arrayVal {
			floatVal, ok := v.(float64)
			if !ok {
				return errors.New("price array must only contain numbers")
			}
			floatVals[idx] = floatVal
		}

		p.amountCents = NewPrice(uint64(floatVals[0]), uint64(floatVals[1])).amountCents
		return nil
	}

	mapVal, isMap := raw.(map[string]interface{})
	if isMap {
		dollars, ok := mapVal["dollars"].(float64)
		if !ok {
			dollars = 0
		}

		cents, ok := mapVal["cents"].(float64)
		if !ok {
			cents = 0
		}

		p.amountCents = NewPrice(uint64(dollars), uint64(cents)).amountCents
		return nil
	}

	return errors.New("unsupported type of price")
}
