package httpexpect

// Boolean provides methods to inspect attached bool value
// (Go representation of JSON boolean).
type Boolean struct {
	chain chain
	value bool
}

// NewBoolean returns a new Boolean given a reporter used to report
// failures and value to be inspected.
//
// reporter should not be nil.
//
// Example:
//  boolean := NewBoolean(t, true)
func NewBoolean(reporter Reporter, value bool) *Boolean {
	return &Boolean{makeChain(reporter), value}
}

// Raw returns underlying value attached to Boolean.
// This is the value originally passed to NewBoolean.
//
// Example:
//  boolean := NewBoolean(t, true)
//  assert.Equal(t, true, boolean.Raw())
func (b *Boolean) Raw() bool {
	return b.value
}

// Equal succeedes if boolean is equal to given value.
//
// Example:
//  boolean := NewBoolean(t, true)
//  boolean.Equal(true)
func (b *Boolean) Equal(value bool) *Boolean {
	if !(b.value == value) {
		b.chain.fail("expected boolean == %v, but got %v", value, b.value)
	}
	return b
}

// NotEqual succeedes if boolean is not equal to given value.
//
// Example:
//  boolean := NewBoolean(t, true)
//  boolean.NotEqual(false)
func (b *Boolean) NotEqual(value bool) *Boolean {
	if !(b.value != value) {
		b.chain.fail("expected boolean != %v, but got %v", value, b.value)
	}
	return b
}

// True succeedes if boolean is true.
//
// Example:
//  boolean := NewBoolean(t, true)
//  boolean.True()
func (b *Boolean) True() *Boolean {
	return b.Equal(true)
}

// False succeedes if boolean is false.
//
// Example:
//  boolean := NewBoolean(t, false)
//  boolean.False()
func (b *Boolean) False() *Boolean {
	return b.Equal(false)
}
