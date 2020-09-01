package pointer

import "time"

// ValueOfBool returns the value of the bool pointer passed in
// or false if the pointer is nil.
func ValueOfBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// ValueOfByte returns the value of the byte pointer passed in
// or 0 if the pointer is nil.
func ValueOfByte(b *byte) byte {
	if b == nil {
		return 0
	}
	return *b
}

// ValueOfComplex64 returns the value of the complex64 pointer passed in
// or 0 if the pointer is nil.
func ValueOfComplex64(c *complex64) complex64 {
	if c == nil {
		return 0
	}
	return *c
}

// ValueOfComplex128 returns the value of the complex128 pointer passed in
// or 0 if the pointer is nil.
func ValueOfComplex128(c *complex128) complex128 {
	if c == nil {
		return 0
	}
	return *c
}

// ValueOfError returns the value of the error pointer passed in
// or nil if the pointer is nil.
func ValueOfError(e *error) error {
	if e == nil {
		return nil
	}
	return *e
}

// ValueOfFloat32 returns the value of the float32 pointer passed in
// or 0 if the pointer is nil.
func ValueOfFloat32(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

// ValueOfFloat64 returns the value of the float64 pointer passed in
// or 0 if the pointer is nil.
func ValueOfFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// ValueOfInt returns the value of the int pointer passed in
// or 0 if the pointer is nil.
func ValueOfInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt8 returns the value of the int8 pointer passed in
// or 0 if the pointer is nil.
func ValueOfInt8(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt16 returns the value of the int16 pointer passed in
// or 0 if the pointer is nil.
func ValueOfInt16(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt32 returns the value of the int32 pointer passed in
// or 0 if the pointer is nil.
func ValueOfInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfInt64 returns the value of the int64 pointer passed in
// or 0 if the pointer is nil.
func ValueOfInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// ValueOfRune returns the value of the rune pointer passed in
// or 0 if the pointer is nil.
func ValueOfRune(r *rune) rune {
	if r == nil {
		return 0
	}
	return *r
}

// ValueOfString returns the value of the string pointer passed in
// or empty string if the pointer is nil.
func ValueOfString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ValueOfUint returns the value of the uint pointer passed in
// or 0 if the pointer is nil.
func ValueOfUint(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint8 returns the value of the uint8 pointer passed in
// or 0 if the pointer is nil.
func ValueOfUint8(u *uint8) uint8 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint16 returns the value of the uint16 pointer passed in
// or 0 if the pointer is nil.
func ValueOfUint16(u *uint16) uint16 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint32 returns the value of the uint32 pointer passed in
// or 0 if the pointer is nil.
func ValueOfUint32(u *uint32) uint32 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUint64 returns the value of the uint64 pointer passed in
// or 0 if the pointer is nil.
func ValueOfUint64(u *uint64) uint64 {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfUintptr returns the value of the uintptr pointer passed in
// or 0 if the pointer is nil.
func ValueOfUintptr(u *uintptr) uintptr {
	if u == nil {
		return 0
	}
	return *u
}

// ValueOfDuration returns the value of the duration pointer passed in
// or 0 if the pointer is nil.
func ValueOfDuration(d *time.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return *d
}

// ValueOfTime returns the value of the time pointer passed in
// or zero time.Time if the pointer is nil.
func ValueOfTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
