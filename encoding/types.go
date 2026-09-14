package encoding

import (
	"encoding/gob"
	"fmt"
	"io"

	"github.com/DNahar74/distributed-file-system/message"
)

type Decoder interface {
	Decode(io.Reader, *message.Message) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *message.Message) error {
	fmt.Println("Hello from GOB decoder")
	err := gob.NewDecoder(r).Decode(msg)
	return err
}

type NOPDecoder struct{}

func (dec NOPDecoder) Decode(r io.Reader, msg *message.Message) error {
	buf := make([]byte, 4096)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	msg.Payload = buf[:n]
	return nil
}

type Encoder interface {
	Encode(io.Writer, any) error
}

type GOBEncoder struct{}

func (enc GOBEncoder) Encode(w io.Writer, v any) error {
	fmt.Println("Hello from Encoder")
	err := gob.NewEncoder(w).Encode(v)
	return err
}
