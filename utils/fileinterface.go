package utils

import (
	"io"
	"os"
)

type MyFileInterface interface {
	Reader() io.Reader
	Name() string
	Close() error
}

// MyFile
type MyFile struct {
	file os.File
}

func NewMyFile(f os.File) MyFile {
	return MyFile{file: f}
}
func (m *MyFile) Reader() io.Reader {
	return &m.file
}
func (m *MyFile) Name() string {
	return m.file.Name()
}
func (m *MyFile) Close() error {
	return m.file.Close()
}

// MyBufferAsFile
type MyBufferAsFile struct {
	buffer io.Reader
	name   string
}

func NewMyBufferAsFile(b io.Reader, n string) MyBufferAsFile {
	return MyBufferAsFile{buffer: b, name: n}
}
func (m *MyBufferAsFile) Reader() io.Reader {
	return m.buffer
}
func (m *MyBufferAsFile) Name() string {
	return m.name
}
func (m *MyBufferAsFile) Close() error {
	return nil
}
