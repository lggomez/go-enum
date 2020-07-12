package internal

import (
	"bytes"
	"go/format"
	"os"
)

func FormatSource(src *bytes.Buffer) ([]byte, error) {
	formattedSrc, err := format.Source(src.Bytes())
	if err != nil {
		return nil, err
	}
	return formattedSrc, nil
}

func SaveFile(path string, formattedSrc []byte) error {
	goFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer goFile.Close()
	_, err = goFile.Write(formattedSrc)
	return err
}
