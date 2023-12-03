package jmap

import (
	"encoding/json"
	"io"
)

// Map and write from object to writer
func Mapto[T any](o T, w io.Writer) error {
	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(o); err != nil {
		return err
	}
	if _, err = w.Write(buf); err != nil {
		return err
	}
	return nil
}

// Map from reader to object
func Mapfrom[T any](r io.Reader, o *T) error {
	var (
		buf []byte
		err error
	)

	if buf, err = io.ReadAll(r); err != nil {
		return err
	}
	if json.Unmarshal(buf, o); err != nil {
		return err
	}
	return nil
}
