package marshal

import (
	"encoding/json"
	"errors"
)

func IsArray(char byte) bool {
	return char == '['
}

func IsString(char byte) bool {
	return char == '"' || char == '\''
}

type Decoder struct {
	Data        []byte
	err         error
	rawMessages []json.RawMessage
}

type ItemDecoder struct {
	Value    any
	Callback func(i int)
}

func NewDecoder(data *[]byte) Decoder {
	return Decoder{Data: *data, err: errors.New("")}
}

func (decoder *Decoder) PrepareArray() int {
	if err := json.Unmarshal(decoder.Data, &decoder.rawMessages); err != nil {
		decoder.err = err
		return 0
	}

	return len(decoder.rawMessages)
}

func (decoder *Decoder) Iterate(itemDecoders ...ItemDecoder) *Decoder {
	for i, msg := range decoder.rawMessages {
		for _, itemDecoder := range itemDecoders {
			decoder.err = json.Unmarshal(msg, &itemDecoder.Value)
			if decoder.err == nil {
				itemDecoder.Callback(i)
				break
			}
		}
	}

	return decoder
}

func (decoder *Decoder) AttemptToDecode(v any, callback func()) *Decoder {
	if decoder.err == nil {
		return decoder
	}

	decoder.err = json.Unmarshal(decoder.Data, &v)
	callback()

	return decoder
}

func (decoder *Decoder) Error() error {
	return decoder.err
}
