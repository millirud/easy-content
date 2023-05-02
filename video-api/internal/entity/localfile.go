package entity

import (
	"io"
	"io/ioutil"
	"os"
)

func NewLocalFileByReader(r io.Reader) (LocalFile, error) {

	file, err := ioutil.TempFile("data", "prefix")

	if err != nil {
		return LocalFile{}, err
	}

	content, err := io.ReadAll(r)

	if err != nil {
		return LocalFile{}, err
	}

	_, err = file.Write(content)

	if err != nil {
		return LocalFile{}, err
	}

	return LocalFile{
		file: file,
	}, nil
}

type LocalFile struct {
	file *os.File
}
