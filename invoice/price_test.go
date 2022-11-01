package invoice

import (
	"encoding/json"
	"testing"
)

func TestNewPrice(t *testing.T) {
	assertAmountCents(t, NewPrice(100, 37), 10037)
}

func TestNewPriceFromFloat(t *testing.T) {
	{
		p, err := NewPriceFromFloat(12.34)
		if err != nil {
			t.Fatal(err)
		}

		expected := NewPrice(12, 34)

		if p != expected {
			t.Fatalf("Got %v, expected: %v", p, expected)
		}
	}
}

func TestNewPriceFromString(t *testing.T) {
	{
		p, err := NewPriceFromString("$1,234,567.89")
		if err != nil {
			t.Fatal(err)
		}

		expected := NewPrice(1234567, 89)

		if p != expected {
			t.Fatalf("Got %v, expected: %v", p, expected)
		}
	}

	{
		p, err := NewPriceFromString("0")
		if err != nil {
			t.Fatal(err)
		}

		expected := NewPrice(0, 0)

		if p != expected {
			t.Fatalf("Got %v, expected: %v", p, expected)
		}
	}

	{
		p, err := NewPriceFromString(".03")
		if err != nil {
			t.Fatal(err)
		}

		expected := NewPrice(0, 3)

		if p != expected {
			t.Fatalf("Got %v, expected: %v", p, expected)
		}
	}

	{
		p, err := NewPriceFromString("$53")
		if err != nil {
			t.Fatal(err)
		}

		expected := NewPrice(53, 0)

		if p != expected {
			t.Fatalf("Got %v, expected: %v", p, expected)
		}
	}

}

func TestPrice_Dollars(t *testing.T) {
	assertDollarAmount(t, NewPrice(100, 0), 100)
	assertDollarAmount(t, NewPrice(0, 37), 0)
	assertDollarAmount(t, NewPrice(10, 23), 10)
}

func TestPrice_Cents(t *testing.T) {
	assertCentsAmount(t, NewPrice(100, 0), 0)
	assertCentsAmount(t, NewPrice(0, 37), 37)
	assertCentsAmount(t, NewPrice(10, 23), 23)
}

func TestPrice_String(t *testing.T) {
	assertAmountString(t, NewPrice(0, 0), "$0.00")
	assertAmountString(t, NewPrice(100, 42), "$100.42")
	assertAmountString(t, NewPrice(5000, 20), "$5,000.20")
	assertAmountString(t, NewPrice(10025, 38), "$10,025.38")
}

func TestPrice_Add(t *testing.T) {
	assertSum(t, NewPrice(100, 0), NewPrice(0, 37), 10037)
	assertSum(t, NewPrice(5, 22), NewPrice(3, 49), 871)
}

func TestPrice_Mul(t *testing.T) {
	assertProduct(t, NewPrice(100, 0), 2, 20000)
	assertProduct(t, NewPrice(9273, 89), 0.198234, 183840)
	assertProduct(t, NewPrice(100, 0), 97.25, 972500)
}

func TestPrice_UnmarshalJSON(t *testing.T) {
	assertJSONEq(t, `[]`, NewPrice(0, 0))
	assertJSONEq(t, `[1]`, NewPrice(1, 0))
	assertJSONEq(t, `[0, 1]`, NewPrice(0, 1))
	assertJSONEq(t, `[12, 34]`, NewPrice(12, 34))
	assertJSONErr(t, `[1,2,3]`)

	assertJSONEq(t, `{}`, NewPrice(0, 0))
	assertJSONEq(t, `{"dollars":1}`, NewPrice(1, 0))
	assertJSONEq(t, `{"cents":1}`, NewPrice(0, 1))
	assertJSONEq(t, `{"cents": 34,"dollars": 12}`, NewPrice(12, 34))

	assertJSONEq(t, `0`, NewPrice(0, 0))
	assertJSONEq(t, `10`, NewPrice(10, 0))
	assertJSONEq(t, `0.0`, NewPrice(0, 0))
	assertJSONEq(t, `12.34`, NewPrice(12, 34))
	assertJSONEq(t, `"0"`, NewPrice(0, 0))
	assertJSONEq(t, `"10"`, NewPrice(10, 0))
	assertJSONEq(t, `"12.34"`, NewPrice(12, 34))
	assertJSONEq(t, `"$0.00"`, NewPrice(0, 0))
	assertJSONEq(t, `"$10.00"`, NewPrice(10, 0))
	assertJSONEq(t, `"$10"`, NewPrice(10, 0))
	assertJSONEq(t, `"$12.34"`, NewPrice(12, 34))
	assertJSONEq(t, `"$1,234,567.89"`, NewPrice(1234567, 89))
}

func assertAmountCents(t *testing.T, price Price, expectedAmountCents uint64) {
	if price.amountCents != expectedAmountCents {
		t.Fatalf(`Got %v, wanted %v`, price.amountCents, expectedAmountCents)
	}
}

func assertDollarAmount(t *testing.T, price Price, expectedDollars uint64) {
	dollars := price.Dollars()

	if dollars != expectedDollars {
		t.Fatalf(`Got %v, wanted %v`, dollars, expectedDollars)
	}
}

func assertCentsAmount(t *testing.T, price Price, expectedCents uint64) {
	cents := price.Cents()

	if cents != expectedCents {
		t.Fatalf(`Got %v, wanted %v`, cents, expectedCents)
	}
}

func assertAmountString(t *testing.T, price Price, expectedString string) {
	amountString := price.String()

	if amountString != expectedString {
		t.Fatalf(`Got %v, wanted %v`, amountString, expectedString)
	}
}

func assertSum(t *testing.T, lhs, rhs Price, expectedAmountCents uint64) {
	result := lhs.Add(rhs)

	if result.amountCents != expectedAmountCents {
		t.Fatalf(`Got %v, wanted %v`, result.amountCents, expectedAmountCents)
	}
}

func assertJSONEq(t *testing.T, jsonStr string, expected Price) {
	var p Price
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		t.Fatal(err)
	}

	if p != expected {
		t.Fatalf("Got %v, wanted %v", p, expected)
	}
}

func assertJSONErr(t *testing.T, jsonStr string) {
	var p Price
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err == nil {
		t.Fatal("Expected error to not be nil")
	}
}

func assertProduct(t *testing.T, lhs Price, rhs float64, expectedAmountCents uint64) {
	result := lhs.Mul(rhs)

	if result.amountCents != expectedAmountCents {
		t.Fatalf(`Got %v, wanted %v`, result.amountCents, expectedAmountCents)
	}
}
