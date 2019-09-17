package codec

// Encoder interface
type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

// Decoder interface
type Decoder interface {
	Decode([]byte, interface{}) error
}

// Codec encode/decode interface
type Codec interface {
	Encoder
	Decoder
}

// Formatter format interface to merge data
type Formatter interface {
	Format(interface{}, interface{}) error
}
