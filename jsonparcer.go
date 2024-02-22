package jsonparcer_go

import (
	"io"
)

func Decode[K any](json io.Reader, output *K) error {
	return nil
}

func Encode[K any](json io.Writer, input *K) error {
	return nil
}
