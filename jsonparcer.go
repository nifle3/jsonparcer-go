package jsonparcer_go

import (
	"io"
)

func Decode[K any](json io.Reader, output *K) error {

	return nil
}

func DecodeJsonToMap(json io.Reader) (map[string]interface{}, error) {
	buffer := make([]byte, 1024)
	result := make(map[string]interface{})

	for {
		n, err := json.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		bufferString := string(buffer[:n])

	}

	return result, nil
}

func Encode[K any](json io.Writer, input *K) error {
	return nil
}
