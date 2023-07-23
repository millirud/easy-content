package entity

import "io"

type File struct {
	reader      io.Reader
	contentType string
	filename    string
	size        int64
}

func NewFile(reader io.Reader, size int64, contentType string, filename string) *File {
	return &File{
		reader:      reader,
		size:        size,
		contentType: contentType,
		filename:    filename,
	}
}

func (f *File) Reader() io.Reader {
	return f.reader
}

func (f *File) ContentType() string {
	return f.contentType
}

func (f *File) Size() int64 {
	return f.size
}

func (f *File) Filename() string {
	return f.filename
}
