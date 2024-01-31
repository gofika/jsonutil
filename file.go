package jsonutil

import (
	"encoding/json"
	"os"

	"github.com/gofika/fileutil"
)

// ReadFileAny read struct from json file
func ReadFileAny(filename string, v any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile read struct from json file
func ReadFile[T any](filename string) (T, error) {
	var v T
	err := ReadFileAny(filename, &v)
	return v, err
}

// WriteFile write struct to json file
func WriteFile(filename string, v any) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := json.NewEncoder(f)
	return enc.Encode(v)
}

// WriteFileIndent write struct to json file with indent
func WriteFileIndent(filename string, v any, indent string) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := json.NewEncoder(f)
	enc.SetIndent("", indent)
	return enc.Encode(v)
}
