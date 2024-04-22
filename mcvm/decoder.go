package mcvm

// A JSON file decoder
type Decoder interface {
	Decode(data []byte) error
	Encode() ([]byte, error)
	DecodeFromFile(path string) error
	EncodeFromFile(path string) error
}
