

// NewBuffer creates and initializes a new Buffer using buf as its
// initial contents.

// The new Buffer takes ownership of buf, and the
// caller should not use buf after this call.

// NewBuffer is intended to prepare a Buffer to read existing data.

// It can also be used to set the initial size of the internal buffer for writing.
// To do that, buf should have the desired capacity but a length of zero.

// In most cases, new(Buffer) (or just declaring a Buffer variable) is
// sufficient to initialize a Buffer.

func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }
