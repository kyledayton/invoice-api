package numberfmt

import "testing"

func assertEq[T comparable](t *testing.T, actual, expected T) {
	if actual != expected {
		t.Fatalf(`Got %v, wanted %v`, actual, expected)
	}
}

func TestSeparated(t *testing.T) {
	assertEq(t, Separated(123456, 3, ","), "123,456")
	assertEq(t, Separated(0, 3, ","), "0")
	assertEq(t, Separated(90210, 2, "^"), "9^02^10")
	assertEq(t, Separated(123, 1, "."), "1.2.3")
}

func TestThousandsSeparated(t *testing.T) {
	assertEq(t, ThousandsSeparated(0, ","), "0")
	assertEq(t, ThousandsSeparated(123, ","), "123")
	assertEq(t, ThousandsSeparated(1002, ","), "1,002")
	assertEq(t, ThousandsSeparated(1234, ","), "1,234")
	assertEq(t, ThousandsSeparated(123456, ","), "123,456")
	assertEq(t, ThousandsSeparated(12345678, ","), "12,345,678")
	assertEq(t, ThousandsSeparated(12345678, "'"), "12'345'678")
	assertEq(t, ThousandsSeparated(82872376232, ","), "82,872,376,232")
}

func TestPow10(t *testing.T) {
	assertEq(t, pow10(0), 1)
	assertEq(t, pow10(1), 10)
	assertEq(t, pow10(2), 100)
	assertEq(t, pow10(3), 1000)
}
