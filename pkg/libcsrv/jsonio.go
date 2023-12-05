package libcsrv

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Read from reader and cast to pointer
func ReadJson(r io.Reader, ptr any) error {
	if buf, err := io.ReadAll(r); err != nil {
		return err
	} else if err = json.Unmarshal(buf, ptr); err != nil {
		return err
	}
	return nil
}

// Cast object to buffer and write to writer
func WriteJson(w io.Writer, obj any) error {
	if buf, err := json.Marshal(obj); err != nil {
		return err
	} else if _, err = fmt.Fprint(w, string(buf)); err != nil {
		return err
	}
	return nil
}

// Cast object to strings reader, return reader
func MakeJsonReader(obj any) (*strings.Reader, error) {
	if buf, err := json.Marshal(obj); err != nil {
		return nil, err
	} else {
		return strings.NewReader(string(buf)), nil
	}
}
