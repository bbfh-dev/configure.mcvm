package mcvm

type Decoder interface {
	Decode(data []byte) error
	Encode() ([]byte, error)
	DecodeFromFile(path string) error
	EncodeFromFile(path string) error
}
