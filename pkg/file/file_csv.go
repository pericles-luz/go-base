package file

import (
	"encoding/csv"
	"os"

	"github.com/pericles-luz/go-base/pkg/utils"
)

type FileCSV struct {
	path   string
	file   *os.File
	reader *csv.Reader
	header []string
}

func NewFileCSV(path string) *FileCSV {
	if !utils.FileExists(path) {
		path = ""
	}
	return &FileCSV{path: path}
}

func (f *FileCSV) Path() string {
	return f.path
}

func (f *FileCSV) Open() error {
	file, err := os.Open(f.Path())
	if err != nil {
		return err
	}
	f.file = file
	f.reader = csv.NewReader(file)
	header, err := f.reader.Read()
	if err != nil {
		return err
	}
	f.header = header
	return nil
}

func (f *FileCSV) Close() error {
	return f.file.Close()
}

func (f *FileCSV) ReadLine() (map[string]string, error) {
	line, err := f.reader.Read()
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	for i, value := range line {
		if i >= len(f.header) {
			break
		}
		data[f.header[i]] = value
	}
	return data, nil
}

func (f *FileCSV) ReadAll() ([]map[string]string, error) {
	var data []map[string]string
	defer f.Close()
	for {
		line, err := f.ReadLine()
		if err != nil {
			return nil, err
		}
		if line == nil {
			break
		}
		data = append(data, line)
	}
	return data, nil
}
