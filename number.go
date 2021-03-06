package httpexpect

// Number provides methods to inspect attached float64 value
// (Go representation of JSON number).
type Number struct {
	chain chain
	value float64
}

// NewNumber returns a new Number given a reporter used to report
// failures and value to be inspected.
//
// reporter should not be nil.
//
// Example:
//  number := NewNumber(t, 123.4)
func NewNumber(reporter Reporter, value float64) *Number {
	return &Number{makeChain(reporter), value}
}

// Raw returns underlying value attached to Number.
// This is the value originally passed to NewNumber.
//
// Example:
//  number := NewNumber(t, 123.4)
//  assert.Equal(t, 123.4, number.Raw())
func (n *Number) Raw() float64 {
	return n.value
}

// Equal succeedes if number is equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.Equal(float64(123))
//  number.Equal(int32(123))
func (n *Number) Equal(value interface{}) *Number {
	v, ok := canonNumber(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value == v) {
		n.chain.fail("expected number == %v, but got %v", v, n.value)
	}
	return n
}

// NotEqual succeedes if number is not equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.NotEqual(float64(321))
//  number.NotEqual(int32(321))
func (n *Number) NotEqual(value interface{}) *Number {
	v, ok := canonNumber(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value != v) {
		n.chain.fail("expected number != %v, but got %v", v, n.value)
	}
	return n
}

// Gt succeedes if number is greater than given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.Gt(float64(122))
//  number.Gt(int32(122))
func (n *Number) Gt(value interface{}) *Number {
	v, ok := canonNumber(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value > v) {
		n.chain.fail("expected number > %v, but got %v", v, n.value)
	}
	return n
}

// Ge succeedes if number is greater than or equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.Ge(float64(122))
//  number.Ge(int32(122))
func (n *Number) Ge(value interface{}) *Number {
	v, ok := canonNumber(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value >= v) {
		n.chain.fail("expected number >= %v, but got %v", v, n.value)
	}
	return n
}

// Lt succeedes if number is lesser than given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.Lt(float64(124))
//  number.Lt(int32(124))
func (n *Number) Lt(value interface{}) *Number {
	v, ok := canonNumber(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value < v) {
		n.chain.fail("expected number < %v, but got %v", v, n.value)
	}
	return n
}

// Le succeedes if number is lesser than or equal to given value.
//
// value should have numeric type convertible to float64. Before comparison,
// it is converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.Le(float64(124))
//  number.Le(int32(124))
func (n *Number) Le(value interface{}) *Number {
	v, ok := canonNumber(&n.chain, value)
	if !ok {
		return n
	}
	if !(n.value <= v) {
		n.chain.fail("expected number <= %v, but got %v", v, n.value)
	}
	return n
}

// InRange succeedes if number is in given range [min; max].
//
// min and max should have numeric type convertible to float64. Before comparison,
// they are converted to float64.
//
// Example:
//  number := NewNumber(t, 123)
//  number.InRange(float32(100), int32(200))  // success
//  number.InRange(100, 200)                  // success
//  number.InRange(123, 123)                  // success
func (n *Number) InRange(min, max interface{}) *Number {
	a, ok := canonNumber(&n.chain, min)
	if !ok {
		return n
	}
	b, ok := canonNumber(&n.chain, max)
	if !ok {
		return n
	}
	if !(n.value >= a && n.value <= b) {
		n.chain.fail("expected number in range [%v; %v], but got %v", a, b, n.value)
	}
	return n
}
